package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	fmt.Println(divisor(a, b, k))
}

func divisor(a, b, k int) int {
	g := gcd(a, b)
	c := 1
	for g > 1 {
		g--
		if a%g == 0 && b%g == 0 {
			c++
			if c == k {
				return g
			}
		}
	}
	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
