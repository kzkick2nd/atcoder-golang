package main

import "fmt"

type pair struct {
	i, j int
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

func grid(h, w int, input []string) string {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == '.' {
				continue
			}
			var count int
			for _, p := range []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				si := i + p.i
				sj := j + p.j
				if si >= 0 &&
					sj >= 0 &&
					si < h &&
					sj < w &&
					input[si][sj] == '#' {
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
