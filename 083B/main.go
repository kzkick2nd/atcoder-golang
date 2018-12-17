package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(sum(n, a, b))
}

func sum(n, a, b int) int {
	var t int
	for i := 1; i <= n; i++ {
		s := sumDigits(i)
		if s >= a && s <= b {
			t += i
		}
	}
	return t
}

func sumDigits(n int) int {
	var s int
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}
