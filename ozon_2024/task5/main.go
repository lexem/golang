package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dir struct {
	Dir     string
	Files   []string
	Folders []Dir
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for k := 0; k < t; k++ {

		var n int
		fmt.Fscanln(in, &n)

		scanner := bufio.NewScanner(os.Stdin)

		var jsonStr string
		for i := 0; i < n; i++ {
			scanner.Scan()
			jsonStr += scanner.Text() + " "
		}

		var dir Dir
		if err := json.Unmarshal([]byte(jsonStr), &dir); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(out, dir)
	}
}
