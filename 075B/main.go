package main

import (
	"fmt"
	"strconv"
)

func main() {
	var f, w int
	fmt.Scan(&f, &w)
	maps := make([]string, f)
	for i := range maps {
		fmt.Scan(&maps[i])
	}
	for _, out := range mine(f, w, maps) {
		fmt.Println(out)
	}
}

func mine(f, w int, maps []string) []string {
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{1, 1, 1, 0, 0, -1, -1, -1}

	minemap := make([]string, f)
	for i := 0; i < f; i++ {
		for j := 0; j < w; j++ {
			if string(maps[i][j]) == "#" {
				minemap[i] += "#"
				continue
			}
			var mine int
			for d := 0; d < 8; d++ {
				posx := i + dx[d]
				posy := j + dy[d]
				if posx < 0 || f <= posx {
					continue
				}
				if posy < 0 || w <= posy {
					continue
				}
				if string(maps[posx][posy]) == "#" {
					mine++
				}
			}
			minemap[i] += strconv.Itoa(mine)
		}
	}
	return minemap
}
