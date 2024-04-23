package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Member struct {
	index int
	time  int
	place int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {

		var n int
		fmt.Fscan(in, &n)

		members := make([]Member, n+1)

		for i := 1; i <= n; i++ {
			var time int
			fmt.Fscan(in, &time)

			members[i] = Member{i, time, -1}
		}

		sort.Slice(members, func(i, j int) bool {
			return members[i].time < members[j].time
		})

		for i, m := range members {
			m.place = i
		}

		sort.Slice(members, func(i, j int) bool {
			return members[i].index < members[j].index
		})

		for _, m := range members {
			fmt.Fprintln(out, m.index, m.time)
		}

		fmt.Fprintln(out)
	}
}
