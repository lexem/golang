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

	var a []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		a = append(a, scanner.Text())
	}
	fmt.Println("Initial array:")
	fmt.Println(strings.Join(a, ", "))
	fmt.Println("**********")

	wordLen := len(a[0])
	for i := wordLen - 1; i >= 0; i-- {
		buskets := make([][]string, 10)

		for j := 0; j < n; j++ {
			digit := getDigit(a[j], i)
			//fmt.Printf("number: %s, digit: %d \n", a[j], digit)

			buskets[digit] = append(buskets[digit], a[j])
		}

		fmt.Printf("Phase %d\n", wordLen-i)
		printPhase(buskets)

		a = basketsToA(buskets)
	}

	fmt.Println("Sorted array:")
	fmt.Println(strings.Join(a, ", "))
}

func printPhase(baskets [][]string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Bucket %d: ", i)
		busket := baskets[i]
		if len(busket) == 0 {
			fmt.Println("empty")
		} else {
			fmt.Println(strings.Join(busket, ", "))
		}
	}
	fmt.Println("**********")
}

func basketsToA(baskets [][]string) []string {
	var newA []string
	for i := 0; i < 10; i++ {
		busket := baskets[i]
		if len(busket) > 0 {
			newA = append(newA, busket...)
		}
	}
	return newA
}

func getDigit(str string, position int) int {
	number := []rune(str)
	digit, _ := strconv.Atoi(string(number[position]))
	return digit
}
