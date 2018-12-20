package main

import "fmt"

func main() {
	in := make([]int, 4)
	for i := range in {
		fmt.Scan(&in[i])
	}
	fmt.Println(robo(in))
}

func robo(in []int) int {
	diff := min(in[1], in[3]) - max(in[0], in[2])
	if diff < 0 {
		return 0
	}
	return diff
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
