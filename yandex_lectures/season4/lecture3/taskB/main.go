package main

import (
	"bufio"
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
	s, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	f, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			scanner.Scan()
			matrix[i][j], _ = strconv.Atoi(scanner.Text())
		}
	}

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

		for i := 1; i <= n; i++ {
			d := matrix[minIndex][i]
			if d == -1 {
				continue
			}

			if minValue+d < dist[i] {
				dist[i] = minValue + d
				previous[i] = minIndex
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
		fmt.Print(result[i], " ")
	}
}
