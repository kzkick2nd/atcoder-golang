package main

func main() {}

func divisor(a, b, k int) int {
	return 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
