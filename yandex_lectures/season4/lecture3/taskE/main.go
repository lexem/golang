package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type City struct {
	time  int
	speed int
}

type Edge struct {
	index int
	dist  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	n := next(scanner)

	var cities = make([]City, n+1)
	for i := 1; i <= n; i++ {
		cities[i] = City{next(scanner), next(scanner)}
	}

	vertex := make(map[int][]Edge)
	for i := 2; i <= n; i++ {
		a := next(scanner)
		b := next(scanner)
		l := next(scanner)

		vertex[a] = append(vertex[a], Edge{b, l})
		vertex[b] = append(vertex[b], Edge{a, l})
	}

	fmt.Print(cities)
}

func next(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}
