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

	var n, q int
	fmt.Fscan(in, &n, &q)

	users := make(map[int]int) // id==0 is last broadcast message

	var t, id int
	var serialNumber int
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &t, &id)

		if t == 1 {
			serialNumber++

			if id == 0 {
				users = make(map[int]int)
			}
			users[id] = serialNumber
		} else {
			message := users[id]
			if message == 0 {
				message = users[0]
			}

			fmt.Fprintln(out, message)
		}
	}

}
