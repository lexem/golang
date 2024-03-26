package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
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
	s := next(scanner)
	f := next(scanner)
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

	visited := make([]bool, n+1)

	minTime := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minTime[i] = math.MaxInt
	}
	minTime[s] = 0

	h := &VillageHeap{Village{s, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		startVillage := heap.Pop(h).(Village)

		if visited[startVillage.index] {
			continue
		}

		for toIndex, schedule := range villages[startVillage.index] {
			for time1, time2 := range schedule {
				if startVillage.time > time1 {
					continue
				}

				if time2 < minTime[toIndex] {
					minTime[toIndex] = time2

					heap.Push(h, Village{toIndex, time2})
				}
			}
		}
		visited[startVillage.index] = true
	}

	if minTime[f] == math.MaxInt {
		fmt.Println("-1")
		return
	}

	fmt.Print(minTime[f])
}

func next(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}
