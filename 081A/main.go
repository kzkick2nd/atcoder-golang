package main

import "fmt"

func main() {
	var in string
	fmt.Scan(&in)
	fmt.Println(one(in))
}

func one(in string) int {
	var i int
	for n := range in {
		if in[n] == '1' {
			i++
		}
	}
	return i
}
