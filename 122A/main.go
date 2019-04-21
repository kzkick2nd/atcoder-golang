package main

import "fmt"

func main() {
	var b string
	fmt.Scan(&b)
	fmt.Println(helix(b))
}

func helix(b string) string {
	if b == "A" {
		return "T"
	}
	if b == "T" {
		return "A"
	}
	if b == "C" {
		return "G"
	}
	if b == "G" {
		return "C"
	}
	return ""
}
