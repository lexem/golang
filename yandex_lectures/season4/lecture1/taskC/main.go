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

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	//fmt.Println(m)

	var b []int
	for i := 0; i < m; i++ {
		scanner.Scan()
		bi, _ := strconv.Atoi(scanner.Text())
		b = append(b, bi)
	}
	//fmt.Println(b)

	c := merge(a, b)
	fmt.Print(strings.Trim(fmt.Sprint(c), "[]"))
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
