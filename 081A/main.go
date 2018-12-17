package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(one(input))
}

func one(input string) int {
	var i int
	for n := range input {
		if input[n] == '1' {
			i++
		}
	}
	return i
}
