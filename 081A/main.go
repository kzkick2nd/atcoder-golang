package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(one(input))
}

func one(input string) int {
	var i int
	if input[0] == '1' {
		i++
	}
	if input[1] == '1' {
		i++
	}
	if input[2] == '1' {
		i++
	}
	return i
}
