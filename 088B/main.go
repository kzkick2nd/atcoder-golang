package main

import (
	"fmt"
	"sort"
)

func game(n int, a []int) int {
	var diff int
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for i, num := range a {
		if i%2 == 0 {
			diff += num
		} else {
			diff -= num
		}
	}
	return diff
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(game(n, a))
}
