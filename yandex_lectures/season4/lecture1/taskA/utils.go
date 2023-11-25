package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return result
}

func ReadInts(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var result []int
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		result = append(result, x)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return result
}
