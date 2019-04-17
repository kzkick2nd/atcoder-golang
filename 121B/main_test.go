package main

import "testing"

func TestSolve(t *testing.T) {
	cases := map[string]struct {
		n, m, c int
		b       []int
		a1      []int
		a2      []int
	}{
		"1": {n: 2, m: 3, c: -10, b: []int{1, 2, 3}, a1: []int{3, 2, 1}, a2: []int{1, 2, 2}},
	}

	for tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := solv()
		})
	}
}
