package main

import (
	"fmt"
)

func dama(N, Y int) (x, y, z int) {
	x, y, z = -1, -1, -1
	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {
			c := N - a - b
			if c < 0 {
				continue
			}
			if Y == 10000*a+5000*b+1000*c {
				fmt.Println(a, b, c)
				x = a
				y = b
				z = c
			}
		}
	}
	return x, y, z
}

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	fmt.Println(dama(n, y))
}
