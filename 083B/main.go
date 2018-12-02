package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	fmt.Println(sum(n, a, b))
}

func sum(n, a, b int) int {
	var total int
	for i := 1; i <= n; i++ {
		sum := sumDigits(i)
		if sum >= a && sum <= b {
			total += i
		}
	}
	return total
}

func sumDigits(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
