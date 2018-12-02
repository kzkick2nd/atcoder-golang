package main

import "fmt"

func change(n int, a []int) int {
	var count int
	odd_flag := false
	for {
		for _, num := range a {
			if num%2 != 0 {
				odd_flag = true
				break
			}
		}
		if odd_flag == true {
			break
		}
		for i, num := range a {
			a[i] = num / 2
		}
		count++
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i, _ := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(change(n, a))
}
