package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	runesLen := len(runes)
	for i := 0; i < runesLen; i++ {
		currentPrefix := runes[0 : i+1]
		prefixLen := len(currentPrefix)
		//fmt.Println(string(currentPrefix), "prefix:", prefixLen)

		for j := 0; j*prefixLen < runesLen; j++ {
			currentSample := runes[j*prefixLen : min(j*prefixLen+prefixLen, runesLen)]
			sampleLen := len(currentSample)
			fmt.Println(string(currentPrefix), string(currentSample))

			check(runes, sampleLen, 0, j*prefixLen)
		}
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
