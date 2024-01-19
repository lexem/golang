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

	// default MaxScanTokenSize 64*1024, but we need 2*10^5
	// maybe we should use bufio.Reader instead
	buf := make([]byte, 0, 256*1024)
	scanner.Buffer(buf, 256*1024)

	scanner.Scan()
	stringer := NewStringer(scanner.Text())

	scanner.Scan()
	requestCount, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < requestCount; i++ {
		scanner.Scan()
		lenS, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		isEqual := stringer.IsEqual(a, b, lenS)
		if isEqual {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
