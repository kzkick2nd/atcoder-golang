package main

import "fmt"

type sushi struct {
	t int
	d int
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	menu := make([]sushi, n)
	for n := range menu {
		fmt.Scan(&menu[n].t, &menu[n].d)
	}
	fmt.Println(k, manzoku(menu))
}

func manzoku(k int, m []sushi) int {
	for i := 0; i < k; i++ {
	}
	return 0
}
