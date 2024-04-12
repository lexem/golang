package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var n int
var a []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	n = next(scanner)

	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = i
	}

	generate(1)
}

func generate(t int) {
	if t == n {
		for i := 1; i <= n; i++ {
			fmt.Print(a[i])
		}
		fmt.Println()
	} else {
		for i := t; i <= n; i++ {
			a[t], a[i] = a[i], a[t]

			generate(t + 1)

			a[t], a[i] = a[i], a[t]
		}
	}
}

func next(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}
