package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(evenodd(a * b))
}

func evenodd(i int) string {
	if i%2 == 0 {
		return "Even"
	}
	return "Odd"
}
