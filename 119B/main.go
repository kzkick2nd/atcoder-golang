package main

import "fmt"

type pair struct {
	x float64
	y string
}

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]pair, n)
	for i := range p {
		fmt.Scan(&p[i].x, &p[i].y)
	}
	fmt.Println(yen(p))
}

func yen(p []pair) float64 {
	var sum float64
	for _, v := range p {
		if v.y == "BTC" {
			sum += v.x * 380000.0
		} else {
			sum += v.x
		}
	}
	return sum
}
