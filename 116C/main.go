package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for n := range h {
		fmt.Scan(&h[n])
	}
	fmt.Println(garden(h))
}

func garden(h []int) int {
	return 0
}
