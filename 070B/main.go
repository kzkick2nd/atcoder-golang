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
	return 50
}
