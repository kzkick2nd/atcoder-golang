package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}
	fmt.Println(poly(n, l))
}

func poly(n int, l []int) string {
	sort.Ints(l)
	max := l[len(l)-1]
	var total int
	for i := range l {
		if i == len(l)-1 {
			continue
		}
		total += l[i]
	}
	if total > max {
		return "Yes"
	}
	return "No"
}
