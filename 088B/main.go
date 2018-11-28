package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var alice, bob int

	fmt.Scan(&n)
	cards := make([]int, n)
	for i := range cards {
		fmt.Scan(&cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	for i, num := range cards {
		if i%2 == 0 {
			alice += num
		} else {
			bob += num
		}
	}
	fmt.Println(alice - bob)
}
