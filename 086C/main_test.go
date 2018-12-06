package main

import "testing"

func TestTravel(t *testing.T) {
	n := 2
	time := []int{3, 6}
	x := []int{1, 1}
	y := []int{2, 1}
	want := "Yes"

	if got := travel(n, time, x, y); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
