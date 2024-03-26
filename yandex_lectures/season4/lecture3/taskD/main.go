package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Schedule map[int]int

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
	next(scanner)
	next(scanner)
	r := next(scanner)

	var villages = make([][]Schedule, n+1)
	for i := 1; i <= n; i++ {
		villages[i] = make([]Schedule, n+1)
	}

	for i := 1; i <= r; i++ {
		from := next(scanner)
		time1 := next(scanner)

		to := next(scanner)
		time2 := next(scanner)

		schedule := villages[from][to]
		if schedule == nil {
			schedule = make(map[int]int)
			villages[from][to] = schedule
		}

		schedule[time1] = time2
	}

}

func next(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}
