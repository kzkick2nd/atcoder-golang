package main

import (
	"fmt"
	"sort"
)

func change(a []int) int {
	var l []int
	for _, n := range a {
		var c int
		for ; n%2 == 0; n /= 2 {
			c++
		}
		l = append(l, c)
	}
	sort.Ints(l)
	return l[0]
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(change(a))
}
