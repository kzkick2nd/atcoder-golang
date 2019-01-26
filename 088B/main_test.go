package main

import "testing"

func TestGame(t *testing.T) {
	cases := map[string]struct {
		n      int
		a      []int
		expect int
	}{
		"1": {n: 2, a: []int{3, 1}, expect: 2},
		"2": {n: 3, a: []int{2, 7, 4}, expect: 5},
		"3": {n: 4, a: []int{20, 18, 2, 18}, expect: 18},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := game(tc.n, tc.a); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
