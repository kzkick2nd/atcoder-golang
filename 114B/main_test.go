package main

import "testing"

func TestDax(t *testing.T) {
	in := 1111111111
	want := 642
	if got := dax(in); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
