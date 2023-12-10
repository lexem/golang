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
	runes := []rune(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		len, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		check(runes, len, a, b)
	}
}

func check(runes []rune, len, a, b int) {
	for i := 0; i < len; i++ {
		if runes[a+i] != runes[b+i] {
			fmt.Println("no")
			return
		}
	}
	fmt.Println("yes")
}
