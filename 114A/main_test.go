package main

import "testing"

func TestIwai(t *testing.T) {
	x := 5
	want := "YES"
	if got := iwai(x); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
