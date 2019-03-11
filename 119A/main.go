package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(heisei(s))
}

func heisei(s string) string {
	s = strings.Replace(s, "/", "", -1)
	i, _ := strconv.Atoi(s)
	if i <= 20190430 {
		return "Heisei"
	}
	return "TBD"
}
