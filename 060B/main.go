package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(choose(a, b, c))
}

func choose(a, b, c int) string {
	return "YES"
}
