package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intsData := read("input.txt")

	num1, num2 := find(intsData[0], intsData[1:])
	fmt.Println(num1, num2)
}

func read(filename string) []int {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stringsData := strings.Split(string(fileBytes), "\n")
	intsData := make([]int, len(stringsData))

	for i, s := range stringsData {
		intsData[i], _ = strconv.Atoi(s)
	}
	return intsData
}

func find(x int, numbers []int) (int, int) {
	var checkedIntsMap = make(map[int]bool)
	for _, num := range numbers {
		if checkedIntsMap[x-num] {
			return x - num, num
		}
		checkedIntsMap[num] = true
	}
	return 0, 0
}
