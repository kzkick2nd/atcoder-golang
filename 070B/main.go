package main

import (
	"fmt"
)

func main() {
	in := make([]int, 4)
	for i := range in {
		fmt.Scan(&in[i])
	}
	fmt.Println(robo(in))
}

func robo(in []int) int {
	return max(0, min(in[1], in[3])-max(in[0], in[2]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
