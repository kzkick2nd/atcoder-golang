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
	var c int
	for i := 0; i < 100; i++ {
		if zeroAll(h) {
			break
		}
		var unZero bool
		for i := range h {
			if h[i] == 0 {
				if unZero {
					c++
				}
				unZero = false
			} else {
				unZero = true
				h[i]--
			}
		}
		if unZero {
			c++
		}
	}
	return c
}

func zeroAll(h []int) bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}
