package main

import "testing"

func TestBurger(t *testing.T) {
	n, x := 50, 4321098765432109
	want := 2160549382716056
	if got := burger(n, x); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
