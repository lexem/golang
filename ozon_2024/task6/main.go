package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Member struct {
	index   int
	maxCard int
	gift    int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, currentCard int
	fmt.Fscan(in, &n, &currentCard)

	members := make([]Member, n)

	for i := 0; i < n; i++ {
		var maxCard int
		fmt.Fscan(in, &maxCard)

		members[i] = Member{i, maxCard, -1}
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].maxCard > members[j].maxCard
	})

	for i := 0; i < n; i++ {
		if members[i].maxCard >= currentCard || currentCard == 0 {
			fmt.Fprintln(out, -1)
			return
		}

		members[i].gift = currentCard
		currentCard--
	}

	sort.Slice(members, func(i, j int) bool {
		return members[i].index < members[j].index
	})

	for i := 0; i < n-1; i++ {
		fmt.Fprint(out, members[i].gift, " ")
	}
	fmt.Fprintln(out, members[n-1].gift)

}
