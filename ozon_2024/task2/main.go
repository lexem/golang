package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var r string
	fmt.Fscan(in, &r)

	sticker := []rune(r)

	var n int
	fmt.Fscan(in, &n)

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b, &r)
		str := []rune(r)

		for j := a; j <= b; j++ {
			sticker[j-1] = str[j-a]
		}
	}

	fmt.Fprintln(out, string(sticker))
}
