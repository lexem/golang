package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	//fmt.Println(n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		ai, _ := strconv.Atoi(scanner.Text())
		a = append(a, ai)
	}
	//fmt.Println(a)

	qSort(0, n)
	fmt.Print(strings.Trim(fmt.Sprint(a), "[]"))
}

func qSort(l, r int) {
	len := r - l
	if len == 0 || len == 1 {
		return
	}

	equal, greater := partition(l, r)

	qSort(l, equal)
	qSort(greater, r)
}

func partition(l, r int) (int, int) {
	xIdx := rand.Intn(r - l)
	x := a[xIdx+l]

	equal := l
	greater := l

	for now := l; now < r; now++ {
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
	}

	return equal, greater
}
