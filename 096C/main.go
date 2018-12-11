package main

import (
	"fmt"
	"os"
)

func grid(h, w int, input []string) string {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	for px := 0; px < h; px++ {
		for py := 0; py < w; py++ {
			if string(input[py][px]) == "." {
				continue
			}
			var count int
			for d := 0; d < 4; d++ {
				posx := px + dx[d]
				posy := py + dy[d]
				if posx < 0 || h <= posx {
					continue
				}
				if posy < 0 || w <= posy {
					continue
				}
				if string(input[posy][posx]) == "#" {
					count++
				}
			}
			if count == 0 {
				return "No"
				os.Exit(0)
			}
		}
	}
	return "Yes"
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	input := make([]string, h)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(grid(h, w, input))
}
