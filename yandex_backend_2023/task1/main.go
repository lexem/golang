package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")

	var commandCount int
	var dcCount int
	var serversCount int

	fmt.Fscanf(file, "%d %d %d\n", &dcCount, &serversCount, &commandCount)

	var dcResetsCount = make([]int, dcCount)
	var serversDisabled = make([][]bool, dcCount)
	for i := range serversDisabled {
		serversDisabled[i] = make([]bool, serversCount)
	}

	for m := 0; m < commandCount; m++ {
		var command string
		var firstNum int
		var secondNum int

		fmt.Fscanf(file, "%s %d %d\n", &command, &firstNum, &secondNum)
		//convert to slice indexes
		firstNum--
		secondNum--

		if command == "RESET" {
			dcResetsCount[firstNum]++

			for i := 0; i < serversCount; i++ { // enable all servers in dc
				serversDisabled[firstNum][i] = false
			}
		}

		if command == "DISABLE" {
			serversDisabled[firstNum][secondNum] = true
		}

		if command == "GETMAX" || command == "GETMIN" {
			var metricSlice = make([]int, dcCount)
			for i := 0; i < dcCount; i++ {
				// find all turned on
				for j := 0; j < serversCount; j++ {
					if !serversDisabled[i][j] {
						metricSlice[i]++
					}
				}
				// multiply to resets count
				metricSlice[i] *= dcResetsCount[i]
			}

			var min, max = findMinAndMaxIndex(metricSlice)

			var result int
			if command == "GETMIN" {
				result = min
			} else {
				result = max
			}

			fmt.Println(result + 1) // plus one because 1-based numeration

		}

	}

}

func findMinAndMaxIndex(a []int) (minIndex int, maxIndex int) {
	var min, max = a[0], a[0]
	for index, value := range a {
		if value < min {
			min = value
			minIndex = index
		}
		if value > max {
			max = value
			maxIndex = index
		}
	}
	return minIndex, maxIndex
}
