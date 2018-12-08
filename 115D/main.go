package main

import "fmt"

func cook(l int) string {
	if l < 1 {
		return "P"
	}
	return "B" + cook(l-1) + "P" + cook(l-1) + "B"
}

func dax(b string, x int) int {
	var ate int
	for i := 0; i < x; i++ {
		if b[i:i+1] == "P" {
			ate++
		}
	}
	return ate
}

func burger(n, x int) int {
	return dax(cook(n), x)
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	fmt.Println(burger(n, x))
}
