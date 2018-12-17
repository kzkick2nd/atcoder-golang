package main

import (
	"fmt"
)

func dama(n, y int) (a, b, c int) {
	d := make([]int, n)
	a, b, c = -1, -1, -1
	for i := range d {
		for j := range d {
			k := n - i - j
			if k < 0 {
				continue
			}
			if y == 10000*j+5000*j+1000*k {
				a = i
				b = j
				c = k
			}
		}
	}
	return a, b, c
}

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	fmt.Println(dama(n, y))
}
