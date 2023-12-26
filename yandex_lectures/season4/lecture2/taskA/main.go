package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const p = 1000000007
const xConst = 257

var x []int
var h []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// default MaxScanTokenSize 64*1024, but we need 2*10^5
	// maybe we should use bufio.Reader instead
	buf := make([]byte, 0, 256*1024)
	scanner.Buffer(buf, 256*1024)

	scanner.Scan()
	initVars(scanner.Text())

	scanner.Scan()
	requestCount, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < requestCount; i++ {
		scanner.Scan()
		lenS, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		isEqual := check(a, b, lenS)
		if isEqual {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func initVars(str string) {
	runes := []rune(" " + str)
	n := len(str)

	x = make([]int, n+1)
	h = make([]int, n+1)
	x[0] = 1

	for i := 1; i < n+1; i++ {
		h[i] = (h[i-1]*xConst + int(runes[i])) % p
		x[i] = (x[i-1] * xConst) % p
	}
}

func check(a, b, len int) bool {
	return (h[a+len]+h[b]*x[len])%p ==
		(h[b+len]+h[a]*x[len])%p
}
