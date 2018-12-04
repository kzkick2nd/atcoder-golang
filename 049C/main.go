package main

import (
	"fmt"
)

func dp(in string) string {
	dp := make([]int, len(in)+1)
	dp[0] = 1
	for i := 0; i < len(in); i++ {
		if dp[i] == 0 {
			continue
		}
		if i+5 <= len(in) && in[i:i+5] == "dream" {
			dp[i+5] = 1
		}
		if i+7 <= len(in) && in[i:i+7] == "dreamer" {
			dp[i+7] = 1
		}
		if i+5 <= len(in) && in[i:i+5] == "erase" {
			dp[i+5] = 1
		}
		if i+6 <= len(in) && in[i:i+6] == "eraser" {
			dp[i+6] = 1
		}
	}
	if dp[len(in)] == 1 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var in string
	fmt.Scan(&in)
	fmt.Println(dp(in))
}
