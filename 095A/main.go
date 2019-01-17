package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(on(s))
}

func on(s string) int {
	t := 700
	for _, v := range s {
		if v == 'o' {
			t += 100
		}
	}
	return t
}
