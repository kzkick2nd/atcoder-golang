package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(xor(n, k, a))
}

func xor(n, k int, a []int) int {
	var s []int
	for i := 1; i <= k; i++ {
		var next bool
		for _, v := range a {
			if v == i {
				next = true
			}
		}
		if next {
			continue
		}
		var t int
		for j := range a {
			fmt.Println(i, a[j], i^a[j])
			t += i ^ a[j]
		}
		fmt.Println(t, "------")
		s = append(s, t)
	}
	sort.Ints(s)
	return s[len(s)-1]
}
