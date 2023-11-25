package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	x, _ := strconv.Atoi(scanner.Text())
	//fmt.Println(x)

	equal := 0
	greater := 0

	for now := 0; now < n; now++ {
		y := a[now]

		if y < x {
			a[now] = a[greater]
			a[greater] = a[equal]
			a[equal] = y

			equal++
			greater++
		}

		if y == x {
			a[now] = a[greater]
			a[greater] = y

			greater++
		}

		//fmt.Println(a)
	}

	fmt.Println(equal)
	fmt.Println(n - equal)
}
