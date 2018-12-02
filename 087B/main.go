package main

import "fmt"

func pattern(a, b, c, x int) int {
	var pattern int
	for aloop := 0; aloop <= a; aloop++ {
		for bloop := 0; bloop <= b; bloop++ {
			for cloop := 0; cloop <= c; cloop++ {
				sum := 500*aloop + 100*bloop + 50*cloop
				if sum == x {
					pattern++
				}
			}
		}
	}
	return pattern
}

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	fmt.Println(pattern(a, b, c, x))
}
