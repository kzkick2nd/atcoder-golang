package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	fmt.Println(divisor(a, b, k))
}

func divisor(a, b, k int) int {
	var i int
	if a > b {
		i = b
	} else {
		i = a
	}
	var count int
	for {
		if a%i == 0 && b%i == 0 {
			count++
		}
		if count == k {
			return i
		}
		i--
	}
}
