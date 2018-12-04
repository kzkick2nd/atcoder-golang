package main

import "fmt"

func bucket(ins []int) int {
	bucket := map[int]int{}
	for i := range ins {
		bucket[ins[i]]++
	}
	fmt.Println(bucket)
	return len(bucket)
}

func main() {
	var in int
	fmt.Scan(&in)
	ins := make([]int, in)
	for i := range ins {
		fmt.Scan(&ins[i])
	}
	fmt.Println(bucket(ins))
}
