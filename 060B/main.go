package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(choose(a, b, c))
}

func choose(a, b, c int) string {
	for i := 1; i <= 100; i++ {
		if i*a%b == c {
			return "YES"
		}
	}
	return "NO"
}
