package main

import "fmt"

func main() {
	var H, W, h, w int
	fmt.Scan(&H, &W, &h, &w)
	fmt.Println(cells(H, W, h, w))
}

func cells(H, W, h, w int) int {
	return (H - h) * (W - w)
}
