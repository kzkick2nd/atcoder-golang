package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	b := make([]int, n)
	for i := range b {
		fmt.Scan(&b[i])
	}
	fmt.Println(seek(n, a, b))
}

func seek(n int, a, b []int) int {
	var s int
	for i := 0; i < n; i++ {
		var sum int
		for _, v := range a[:i+1] {
			sum += v
		}
		for _, v := range b[i:] {
			sum += v
		}
		if sum > s {
			s = sum
		}
	}
	return s
}
