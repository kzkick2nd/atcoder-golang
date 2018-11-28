package main

import "fmt"

func main() {
	var n, a, b, total int
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n; i++ {
		sum := sumDigits(i)
		if sum >= a && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}

func sumDigits(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
