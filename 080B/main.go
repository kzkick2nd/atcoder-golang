package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(harshad(n))
}

func harshad(n int) string {
	f := digitSum(n)
	if n%f == 0 {
		return "Yes"
	}
	return "No"
}

func digitSum(n int) int {
	var s int
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}
