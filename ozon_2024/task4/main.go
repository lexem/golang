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

	for k := 0; k < t; k++ {

		var n int
		fmt.Fscan(in, &n)

		members := make([]Member, n)

		for i := 0; i < n; i++ {
			var time int
			fmt.Fscan(in, &time)

			members[i] = Member{i, time, -1}
		}

		sort.Slice(members, func(i, j int) bool {
			return members[i].time < members[j].time
		})

		for i := 0; i < n; i++ {
			currentPlace := i + 1
			members[i].place = currentPlace

			for j := i + 1; j < n; j++ {
				if members[j].time-members[j-1].time <= 1 {
					members[j].place = currentPlace
					i++
				} else {
					break
				}
			}
		}
		//fmt.Fprintln(out, members)

		sort.Slice(members, func(i, j int) bool {
			return members[i].index < members[j].index
		})

		for i := 0; i < n-1; i++ {
			fmt.Fprint(out, members[i].place, " ")
		}
		fmt.Fprintln(out, members[n-1].place)
	}
}
