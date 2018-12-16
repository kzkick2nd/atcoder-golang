package main

import "fmt"

func coins(a, b, c, x int) int {
	var cnt int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	fmt.Println(coins(a, b, c, x))
}
