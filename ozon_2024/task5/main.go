package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Dir struct {
	Dir     string
	Files   []string
	Folders []Dir
}

func main() {
	decoder := json.NewDecoder(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	decoder.Decode(&t)

	for k := 0; k < t; k++ {
		var n int
		decoder.Decode(&n) // read just for skipping, we don't need it

		var dir Dir
		decoder.Decode(&dir)

		fmt.Fprintln(out, countHackedFiles(dir, false))
	}
}

func countHackedFiles(dir Dir, infected bool) (result int) {
	infected = infected || isDirInfected(dir) // isDirInfected will not calculated if infected==true already

	if infected {
		result += len(dir.Files)
	}

	for _, folder := range dir.Folders {
		result += countHackedFiles(folder, infected)
	}

	return
}

func isDirInfected(dir Dir) bool {
	for _, fileName := range dir.Files {
		if strings.HasSuffix(fileName, ".hack") {
			return true
		}
	}
	return false
}
