package main

import "testing"

func TestGame(t *testing.T) {
	n := 3
	a := []int{2, 7, 4}
	want := 5
	if got := game(n, a); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
