package main

import (
	"fmt"
)

func dream(s string) string {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := range s {
		if dp[i] == 0 {
			continue
		}
		if i+5 <= len(s) && s[i:i+5] == "dream" {
			dp[i+5] = 1
		}
		if i+7 <= len(s) && s[i:i+7] == "dreamer" {
			dp[i+7] = 1
		}
		if i+5 <= len(s) && s[i:i+5] == "erase" {
			dp[i+5] = 1
		}
		if i+6 <= len(s) && s[i:i+6] == "eraser" {
			dp[i+6] = 1
		}
	}
	if dp[len(s)] == 1 {
		return "YES"
	}
	return "NO"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(dream(s))
}
