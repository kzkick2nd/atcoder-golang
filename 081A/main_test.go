package main

import "testing"

func TestOne(t *testing.T) {
	in := "101"
	want := 2
	if got := one(in); got != want {
		t.Fatalf("want = %v, got = %v", want, got)
	}
}
