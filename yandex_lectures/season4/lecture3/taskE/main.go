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

type City struct {
	time  int
	speed int
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

	// find route from moscow to other cities
	minTime, _ := route(vertex, cities, 1)

	// try optimize minTime by dijkstra principle
	visited := make([]bool, n+1)

	h := &EdgeHeap{}
	heap.Init(h)
	for i := 2; i <= n; i++ {
		heap.Push(h, Edge{i, minTime[i]})
	}

	for h.Len() > 0 {
		bestCityIdx := heap.Pop(h).(Edge).index
		if visited[bestCityIdx] {
			continue
		}

		toBestCityTime, _ := route(vertex, cities, bestCityIdx)

		for i := 2; i <= n; i++ {
			updatedTime := toBestCityTime[i] + minTime[bestCityIdx]
			if updatedTime < minTime[i] {
				minTime[i] = updatedTime
				heap.Push(h, Edge{i, updatedTime})
			}
		}

		visited[bestCityIdx] = true
	}

	fmt.Println(minTime)
}

func findDist(edges []Edge, index int) int {
	for _, edge := range edges {
		if edge.index == index {
			return edge.dist
		}
	}
	return -1
}

func route(vertex map[int][]Edge, cities []City, finishIdx int) (minTime []int, previous []int) {
	n := len(cities) - 1

	visited := make([]bool, n+1)
	previous = make([]int, n+1)

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt
	}
	dist[finishIdx] = 0

	h := &EdgeHeap{Edge{finishIdx, 0}}
	heap.Init(h)

	for h.Len() > 0 {
		startEdge := heap.Pop(h).(Edge)

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
	//fmt.Println(previous)

	// calculate time for each member without changes vehicle
	minTime = make([]int, n+1)

	for i := 2; i <= n; i++ { //exclude moscow, start from 2
		s := finishIdx
		f := i

		time := cities[i].time
		speed := cities[i].speed

		for j := f; j != s; j = previous[j] {
			from := j
			to := previous[j]
			time += findDist(vertex[from], to) / speed
		}

		minTime[i] = time
	}
	//fmt.Println(minTime)

	return
}

func next(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}
