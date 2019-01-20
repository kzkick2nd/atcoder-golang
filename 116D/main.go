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
	fmt.Println(manzoku(menu))
}

func manzoku(m []sushi) int {
	return 0
}
