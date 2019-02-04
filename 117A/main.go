package main

import "fmt"

func main() {
	var t, x float64
	fmt.Scan(&t, &x)
	fmt.Println(leap(t, x))
}

func leap(t, x float64) float64 {
	return t / x
}
