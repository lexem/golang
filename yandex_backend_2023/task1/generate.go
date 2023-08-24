package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {

	var dcCount = 3
	var serversCount = 3

	var commandCount = 12

	f, _ := os.Create("generated.txt")
	defer f.Close()

	f.WriteString(fmt.Sprintf("%d %d %d", dcCount, serversCount, commandCount))

	for i := 0; i < commandCount; i++ {
		f.WriteString("\n")

		switch rand.Intn(4) {
		case 0:
			f.WriteString(fmt.Sprintf("DISABLE %d %d", rand.Intn(dcCount)+1, rand.Intn(serversCount)+1))
		case 1:
			f.WriteString(fmt.Sprintf("RESET %d", rand.Intn(dcCount)+1))
		case 2:
			f.WriteString("GETMIN")
		case 3:
			f.WriteString("GETMAX")

		default:
		}

	}

}
