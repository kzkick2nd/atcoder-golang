package main

import "fmt"

func grid(h, w int, input []string) string {
	for px := 0; px < h; px++ {
		for py := 0; py < w; py++ {
			if string(input[px][py]) == "." {
				continue
			}
			var count int
			if py-1 > 0 {
				if string(input[px][py-1]) == "#" {
					count++
				}
			}
			if px-1 > 0 {
				if string(input[px-1][py]) == "#" {
					count++
				}
			}
			if px+1 < h {
				if string(input[px+1][py]) == "#" {
					count++
				}
			}
			if py+1 < w {
				if string(input[px][py+1]) == "#" {
					count++
				}
			}
			if count == 0 {
				return "No"
			}
		}
	}
	return "Yes"
}

func main() {
	var h, w int
	input := make([]string, h)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(grid(h, w, input))
}
