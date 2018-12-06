package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func travel(n int, t, x, y []int) string {
	post, posx, posy := 0, 0, 0
	for i, _ := range t {
		difft := abs(t[i] - post)
		diffx := abs(x[i] - posx)
		diffy := abs(y[i] - posy)
		if difft < diffx+diffy {
			return "No"
		}
		if even(difft) != even(diffx+diffy) {
			return "No"
		}
		post = t[i]
		posx = x[i]
		posy = y[i]
	}
	return "Yes"
}

func main() {
	var n int
	fmt.Scan(&n)
	time := make([]int, n)
	x := make([]int, n)
	y := make([]int, n)
	for i := range time {
		fmt.Scan(&time[i])
		fmt.Scan(&x[i])
		fmt.Scan(&y[i])
	}
	fmt.Println(travel(n, time, x, y))
}
