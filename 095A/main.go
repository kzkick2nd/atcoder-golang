package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(on(s))
}

func on(s string) int {
	return 700 + strings.Count(s, "o")*100
}
