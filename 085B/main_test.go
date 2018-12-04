package main

import "testing"

func TestMochi(t *testing.T) {
	in := []int{1, 7, 5, 5, 7, 3, 3}
	want := 4
	if got := bucket(in); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
