package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(favSound(a, b, c))
}

func favSound(a, b, c int) int {
	if b/a > c {
		return c
	}
	return b / a
}
