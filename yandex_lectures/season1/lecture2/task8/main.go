package main

import (
	"fmt"
	"os"
)

func main() {
	str := read("input.txt")
	runes := []rune(str)

	lastRune := runes[0]
	lastIndex := 0

	for i, r := range runes {
		if r != lastRune {
			fmt.Printf("%c%d", lastRune, i-lastIndex)
			lastRune = r
			lastIndex = i
		}
	}

	fmt.Printf("%c%d", lastRune, len(runes)-lastIndex)
}

func read(filename string) string {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(fileBytes)
}
