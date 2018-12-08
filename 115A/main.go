package main

import (
	"fmt"
	"strings"
)

func eve(d int) string {
	return "Christmas" + strings.Repeat(" Eve", 25-d)
}

func main() {
	var d int
	fmt.Scan(&d)
	fmt.Println(eve(d))
}
