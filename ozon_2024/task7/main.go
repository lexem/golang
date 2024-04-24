package main

import (
	"bufio"
	"fmt"
	"os"
)

type Member struct {
	index  int
	window int
	move   int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for k := 0; k < t; k++ {

		var n, m int
		fmt.Fscan(in, &n, &m)

		members := make([]Member, m)

		for i := 0; i < m; i++ {
			var w int
			fmt.Fscan(in, &w)

			members[i] = Member{i + 1, w, 0}
		}
		fmt.Fprintln(out, members)
	}
}
