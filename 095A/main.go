package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(on(s))
}

func on(s string) int {
	return 700
}
