package main

import "testing"

func TestEvenodd(t *testing.T) {
	a, b := 3, 4
	want := "Even"
	if got := evenodd(a, b); got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}
