package main

import "testing"

func TestEve(t *testing.T) {
	d := 22
	want := "Christmas Eve Eve Eve"
	if got := eve(d); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
