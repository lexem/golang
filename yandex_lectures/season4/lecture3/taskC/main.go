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

		vertex[a] = append(vertex[a], Edge{index: b, dist: l})
		vertex[b] = append(vertex[b], Edge{index: a, dist: l})
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

	h := &EdgeHeap{Edge{s, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		startEdge := heap.Pop(h).(Edge) //(*h)[0]

		if visited[startEdge.index] {
			continue
		}

		for _, edge := range vertex[startEdge.index] {
			if startEdge.dist+edge.dist < dist[edge.index] {
				dist[edge.index] = startEdge.dist + edge.dist
				previous[edge.index] = startEdge.index

				heap.Push(h, Edge{edge.index, startEdge.dist + edge.dist})
			}
		}
		visited[startEdge.index] = true
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
