package main

import "fmt"

func bucket(d []int) int {
	b := map[int]int{}
	for i := range d {
		b[d[i]]++
	}
	return len(b)
}

func main() {
	var n int
	fmt.Scan(&n)
	d := make([]int, n)
	for i := range d {
		fmt.Scan(&d[i])
	}
	fmt.Println(bucket(d))
}
