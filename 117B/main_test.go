package main

import "testing"

func TestPoly(t *testing.T) {
	cases := map[string]struct {
		n      int
		l      []int
		expect string
	}{
		"1": {n: 4, l: []int{3, 8, 5, 1}, expect: "Yes"},
		"2": {n: 4, l: []int{3, 8, 4, 1}, expect: "No"},
		"3": {n: 10, l: []int{1, 8, 10, 5, 8, 12, 34, 100, 11, 3}, expect: "No"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := poly(tc.n, tc.l); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
