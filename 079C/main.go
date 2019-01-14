package main

import (
	"fmt"
)

func main() {
	var in string
	fmt.Scan(&in)
	fmt.Println(op(in))
}

func op(in string) string {
	nums := [4]int{int(in[0] - '0'), int(in[1] - '0'), int(in[2] - '0'), int(in[3] - '0')}
	for i := 0; i < (1 << 3); i++ {
		ans := nums[0]
		ops := []string{}
		for j := uint8(0); j < 3; j++ {
			if i&(1<<j) != 0 {
				ans += nums[j+1]
				ops = append(ops, "+")
			} else {
				ans -= nums[j+1]
				ops = append(ops, "-")
			}
		}
		if ans == 7 {
			return fmt.Sprintf("%d%s%d%s%d%s%d=7", nums[0], ops[0], nums[1], ops[1], nums[2], ops[2], nums[3])
		}
	}
	return ""
}
