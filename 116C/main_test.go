package main

import "testing"

func TestGarden(t *testing.T) {
	cases := map[string]struct {
		n      int
		h      []int
		expect int
	}{
		"1": {n: 4, h: []int{1, 2, 2, 1}, expect: 2},
		"2": {n: 5, h: []int{3, 1, 2, 3, 1}, expect: 5},
		"3": {n: 8, h: []int{4, 23, 75, 0, 23, 96, 50, 100}, expect: 221},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := garden(tc.h); actual != tc.expect {
				t.Fatalf("expect is %v, got is %v", tc.expect, actual)
			}
		})
	}
}
