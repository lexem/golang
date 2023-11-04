package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type obj struct {
	value int
	index int
}

type sortObj []obj

func (word sortObj) Less(i, j int) bool {
	return word[i].value < word[j].value
}

func (word sortObj) Swap(i, j int) {
	word[i], word[j] = word[j], word[i]
}

func (word sortObj) Len() int {
	return len(word)
}

func main() {
	nums, _ := ReadInts("input.txt")

	_, m := nums[0], nums[1]
	//fmt.Println(n, m)

	prices := nums[2:]
	sort.Sort(sortObj(prices))

	min := prices[0].index
	max := prices[0].index
	fmt.Println(prices)
	for _, p := range prices[:m.value] {
		if p.index < min {
			min = p.index
		}

		if p.index > max {
			max = p.index
		}
	}

	fmt.Println(max - min + 1)
}

func ReadInts(filename string) ([]obj, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	var result []obj

	i := 0
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, obj{value: x, index: i})
		i++
	}
	return result, scanner.Err()
}
