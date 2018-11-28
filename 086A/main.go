package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	x := a * b
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
