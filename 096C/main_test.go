package main

import "testing"

func TestGrid(t *testing.T) {
	h, w := 5, 3
	input := []string{"#.#.#", ".#.#.", "#.#.#"}
	if expected, actual := "No", grid(h, w, input); expected != actual {
		t.Errorf("wont %s but got %s", expected, actual)
	}
}
