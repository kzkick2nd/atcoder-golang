package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(num(n))
}

func num(n int) int {
	var m int
	a := 1
	for ; n > 0; n-- {
		var c int
		for t := n; t%2 == 0; t /= 2 {
			c++
		}
		if c > m {
			m = c
			a = n
		}
	}
	return a
}
