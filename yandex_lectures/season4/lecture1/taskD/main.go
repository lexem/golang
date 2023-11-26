package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var a []int
var merged []int

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

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		ai, _ := strconv.Atoi(scanner.Text())
		a = append(a, ai)
	}
	//fmt.Println(a)

	merged = make([]int, n)
	mergeSort(0, n)
	fmt.Print(strings.Trim(fmt.Sprint(a), "[]"))
}

func mergeSort(lIdx int, rIdx int) {
	lenA := rIdx - lIdx
	if lenA == 1 {
		return
	}

	centerIdx := lenA/2 + lIdx
	mergeSort(lIdx, centerIdx)
	mergeSort(centerIdx, rIdx)

	merge(lIdx, centerIdx, rIdx)
}

func merge(lIdx, centerIdx, rIdx int) {
	lenA := centerIdx - lIdx
	lenB := rIdx - centerIdx

	aPointer := lIdx
	bPointer := centerIdx

	for i := lIdx; i < lIdx+lenA+lenB; i++ {
		if aPointer == lIdx+lenA {
			merged[i] = a[bPointer]
			bPointer++
			continue
		}

		if bPointer == centerIdx+lenB {
			merged[i] = a[aPointer]
			aPointer++
			continue
		}

		if a[aPointer] < a[bPointer] {
			merged[i] = a[aPointer]
			aPointer++
		} else {
			merged[i] = a[bPointer]
			bPointer++
		}
	}

	for i := lIdx; i < lIdx+lenA+lenB; i++ {
		a[i] = merged[i]
	}
}
