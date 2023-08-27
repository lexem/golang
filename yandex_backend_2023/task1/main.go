package main

import (
	"fmt"
	"os"
	"sort"
)

type DataCenter struct {
	index            int
	resetCount       int
	turnedOffServers map[int]bool
}

func (dc DataCenter) metric(serversCount int) int {
	turnedOn := serversCount - len(dc.turnedOffServers)
	return turnedOn * dc.resetCount
}

func main() {

	file, _ := os.Open("input.txt")

	var commandCount int
	var dcCount int
	var serversCount int

	fmt.Fscanf(file, "%d %d %d\n", &dcCount, &serversCount, &commandCount)

	var dcMap = make(map[int]DataCenter)

	for m := 0; m < commandCount; m++ {
		var command string
		var dcIndex int
		var sIdx int

		fmt.Fscanf(file, "%s %d %d\n", &command, &dcIndex, &sIdx)

		if command == "RESET" {
			if dc, ok := dcMap[dcIndex]; !ok {
				dc = DataCenter{
					index:            dcIndex,
					resetCount:       0,
					turnedOffServers: nil,
				}
				dcMap[dcIndex] = dc
			}

			dc := dcMap[dcIndex]
			dc.resetCount++
			dc.turnedOffServers = nil
			dcMap[dcIndex] = dc
		}

		if command == "DISABLE" {
			if dc, ok := dcMap[dcIndex]; !ok {
				dc = DataCenter{
					index:            dcIndex,
					resetCount:       0,
					turnedOffServers: make(map[int]bool),
				}
				dcMap[dcIndex] = dc
			} else if dc.turnedOffServers == nil {
				dc.turnedOffServers = make(map[int]bool)
				dcMap[dcIndex] = dc
			}

			dc := dcMap[dcIndex]
			dc.turnedOffServers[sIdx] = true
		}

		if command == "GETMAX" || command == "GETMIN" {
			var minIdx, maxIdx = findMinAndMaxIndex(dcMap, serversCount)

			if command == "GETMIN" {
				fmt.Println(minIdx)
			} else {
				fmt.Println(maxIdx)
			}

		}

	}

}

func findMinAndMaxIndex(dcMap map[int]DataCenter, serversCount int) (minIndex int, maxIndex int) {
	keys := sortedKeys(dcMap)
	//fmt.Println("--------", keys)
	firstKey := keys[0]

	var min = dcMap[firstKey].metric(serversCount)
	var max = min

	minIndex = dcMap[firstKey].index
	maxIndex = minIndex

	for _, key := range keys {
		dc := dcMap[key]
		metric := dc.metric(serversCount)
		//fmt.Println("--------", metric, dc)
		if metric < min {
			min = metric
			minIndex = dc.index
		}
		if metric > max {
			max = metric
			maxIndex = dc.index
		}
	}

	return minIndex, maxIndex
}

func sortedKeys(m map[int]DataCenter) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	return keys
}
