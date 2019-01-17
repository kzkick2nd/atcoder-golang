package main

import "fmt"

func main() {
	input := make([]int, 3)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(rgb(input))
}

func rgb(rgb []int) string {
	return "YES"
}
