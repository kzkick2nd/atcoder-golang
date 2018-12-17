package main

import "fmt"

func change(n int, list []int) int {
	var c int
	var oddFlag bool
	for {
		for i, num := range list {
			if num%2 != 0 {
				oddFlag = true
				break
			}
			list[i] = num / 2
		}
		if oddFlag == true {
			break
		}
		c++
	}
	return c
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	fmt.Println(change(n, a))
}
