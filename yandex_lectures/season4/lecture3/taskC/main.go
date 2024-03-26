package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Edge struct {
	to  int
	len int
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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	var vertex = make(map[int][]Edge)

	for i := 1; i <= k; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		l, _ := strconv.Atoi(scanner.Text())

		vertex[a] = append(vertex[a], Edge{to: b, len: l})
		vertex[b] = append(vertex[b], Edge{to: a, len: l})
	}

	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	f, _ := strconv.Atoi(scanner.Text())

	visited := make([]bool, n+1)
	previous := make([]int, n+1)

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt
	}
	dist[s] = 0

	for {
		minIndex := -1
		minValue := math.MaxInt

		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}

			if dist[i] < minValue {
				minIndex = i
				minValue = dist[i]
			}
		}

		if minIndex == -1 {
			break
		}

		for _, edge := range vertex[minIndex] {
			if minValue+edge.len < dist[edge.to] {
				dist[edge.to] = minValue + edge.len
				previous[edge.to] = minIndex
			}
		}
		visited[minIndex] = true
	}

	if dist[f] == math.MaxInt {
		fmt.Println("-1")
		return
	}

	result := make([]int, 0)
	result = append(result, f)

	for i := f; i != s; i = previous[i] {
		result = append(result, previous[i])
	}

	for i := len(result) - 1; i >= 0; i-- {
		// path
		//fmt.Print(result[i], " ")
	}

	fmt.Print(dist[f])
}
