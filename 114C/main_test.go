package main

import "testing"

func TestTarget(t *testing.T) {
	in := 575
	want := 4
	if got := target(in); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
