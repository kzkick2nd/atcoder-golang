package main

import "testing"

func TestGrid(t *testing.T) {
	h, w := 3, 3
	input := []string{".#.", "###", ".#."}
	if expected, actual := "Yes", grid(h, w, input); expected != actual {
		t.Errorf("wont %s but got %s", expected, actual)
	}
}
