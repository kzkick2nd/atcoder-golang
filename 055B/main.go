package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(train(n))
}

func train(n int) int {
	w, p := int(math.Pow10(9.0))+7, 1
	for i := 1; i <= n; i++ {
		p = p * i % w
	}
	return p
}
