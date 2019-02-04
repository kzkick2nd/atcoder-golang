package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	x := make([]int, m)
	for i := range x {
		fmt.Scan(&x[i])
	}
	fmt.Println(stream(n, m, x))
}

func stream(n, m int, x []int) int {
	sort.Ints(x)
	var d []int
	for i := range x {
		if i+1 == len(x) {
			continue
		}
		d = append(d, x[i+1]-x[i])
	}
	sort.Ints(d)
	var t int
	for i := range d {
		if i+1 > len(d)-(n-1) {
			continue
		}
		t += d[i]
	}
	return t
}
