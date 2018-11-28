package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)
	s := make([]int, n)
	for i, _ := range s {
		fmt.Scan(&s[i])
	}
	flag := true
	for {
		for _, num := range s {
			if num%2 != 0 {
				flag = false
				break
			}
		}
		if flag == false {
			break
		}
		for i, num := range s {
			s[i] = num / 2
		}
		c++
	}
	fmt.Println(c)
}
