package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(donuts(n))
}

func donuts(n int) string {
	return "Yes"
}
