package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	//fmt.Println(n)

	var a []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		ai, _ := strconv.Atoi(scanner.Text())
		a = append(a, ai)
	}
	//fmt.Println(a)

	a = mergeSort(a)
	fmt.Print(strings.Trim(fmt.Sprint(a), "[]"))
}

func mergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	centerIdx := len(a) / 2
	left := mergeSort(a[:centerIdx])
	right := mergeSort(a[centerIdx:])

	return merge(left, right)
}

func merge(a, b []int) []int {
	lenA := len(a)
	lenB := len(b)

	lenC := lenA + lenB
	c := make([]int, lenC)

	aPointer := 0
	bPointer := 0

	for i := 0; i < lenC; i++ {
		if aPointer == lenA {
			c[i] = b[bPointer]
			bPointer++
			continue
		}

		if bPointer == lenB {
			c[i] = a[aPointer]
			aPointer++
			continue
		}

		if a[aPointer] < b[bPointer] {
			c[i] = a[aPointer]
			aPointer++
		} else {
			c[i] = b[bPointer]
			bPointer++
		}
	}

	return c
}
