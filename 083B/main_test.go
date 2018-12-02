package main

import "testing"

func TestSum(t *testing.T) {
	n, a, b := 20, 2, 5
	want := 84
	if got := sum(n, a, b); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
