package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(evenodd(a, b))
}

func evenodd(a, b int) string {
	x := a * b
	if x%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
