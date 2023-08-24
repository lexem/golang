package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const blockSym = `█`

	file, _ := os.Open("input.txt")

	var line int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &line) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, line)
	}

	var max = 0
	var maxIdx = 0

	for i, val := range nums {
		if max < val {
			max = val
			maxIdx = i
		}
	}

	fmt.Println("max:", max)
	fmt.Println("maxIdx", maxIdx)

	for j := max; j > 0; j-- {
		for i := 0; i < len(nums); i++ {
			if j <= nums[i] {
				fmt.Print(blockSym + blockSym)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	var waterCount = 0
	var localMax = 0
	for i := 0; i < maxIdx; i++ {
		if nums[i] > localMax {
			localMax = nums[i]
		}
		waterCount += localMax - nums[i]
	}

	localMax = 0
	for i := len(nums) - 1; i > maxIdx; i-- {
		if nums[i] > localMax {
			localMax = nums[i]
		}
		waterCount += localMax - nums[i]
	}

	fmt.Println("waterCount:", waterCount)
}
