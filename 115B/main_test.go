package main

import "testing"

func TestSale(t *testing.T) {
	n := 4
	prices := []int{4320, 4320, 4320, 4320}
	want := 15120
	if got := sale(n, prices); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
