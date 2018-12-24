package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(paintBall(n, k))
}

func paintBall(n, k int) int {
	pat := float64(k) * math.Pow(float64(k-1), float64(n-1))
	return int(pat)
}
