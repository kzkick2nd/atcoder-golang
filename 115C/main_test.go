package main

import "testing"

func TestTree(t *testing.T) {
	n, k := 5, 3
	h := []int{5, 7, 5, 7, 7}
	want := 0
	if got := tree(n, k, h); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
