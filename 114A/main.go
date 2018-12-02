package main

import "fmt"

func iwai(x int) string {
	if x == 3 || x == 5 || x == 7 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(iwai(x))
}
