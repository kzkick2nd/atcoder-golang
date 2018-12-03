package main

import "testing"

func TestTarget(t *testing.T) {
	in := 999999999
	want := 26484
	if got := target(in); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
