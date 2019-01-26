package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	l := make([]int, n)
	for i := range l {
		fmt.Scan(&l[i])
	}
	fmt.Println(sunuke(k, l))
}

func sunuke(k int, l []int) int {
	var s int
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	for i := 0; i < k; i++ {
		s += l[i]
	}
	return s
}
