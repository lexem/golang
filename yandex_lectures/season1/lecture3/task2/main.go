package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dictionary := read("dictionary.txt")
	text := read("text.txt")

	result := check(dictionary, text)
	fmt.Println(result)
}

func check(dictionary []string, text []string) []string {
	dictionaryExtend := make(map[string]bool)
	for _, word := range dictionary {
		for i := 0; i < len(word); i++ {
			newWord := delChar(word, i)
			dictionaryExtend[newWord] = true
		}
	}

	result := make([]string, 0)
	for _, word := range text {
		if dictionaryExtend[word] {
			result = append(result, word)
		}
	}

	return result
}

func delChar(str string, index int) string {
	s := []rune(str)
	result := append(s[0:index], s[index+1:]...)
	return string(result)
}

func read(filename string) []string {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(fileBytes), "\n")
}
