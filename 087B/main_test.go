package main

import "testing"

func TestPattern(t *testing.T) {
	a, b, c, x := 2, 2, 2, 100
	want := 2
	if got := pattern(a, b, c, x); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
