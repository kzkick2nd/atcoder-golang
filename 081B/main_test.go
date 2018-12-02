package main

import "testing"

func TestChange(t *testing.T) {
	n := 3
	a := []int{8, 12, 40}
	want := 2
	if got := change(n, a); got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
