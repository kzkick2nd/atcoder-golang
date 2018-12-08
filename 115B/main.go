package main

import (
	"fmt"
	"sort"
)

func sale(n int, prices []int) int {
	sort.Ints(prices)
	subtotal := prices[len(prices)-1] / 2
	for _, v := range prices[:len(prices)-1] {
		subtotal += v
	}
	return subtotal
}

func main() {
	var n int
	fmt.Scan(&n)
	prices := make([]int, n)
	for i := range prices {
		fmt.Scan(&prices[i])
	}
	fmt.Println(sale(n, prices))
}
