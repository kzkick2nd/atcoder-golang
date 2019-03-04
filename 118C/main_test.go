package main

import "testing"

func TestMonst(t *testing.T) {
	cases := map[string]struct {
		n      int
		a      []int
		expect int
	}{
		"1": {n: 4, a: []int{2, 10, 8, 40}, expect: 2},
		"2": {n: 4, a: []int{5, 13, 8, 1000000000}, expect: 1},
		"3": {n: 3, a: []int{1000000000, 1000000000, 1000000000}, expect: 1},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := monster(tc.a); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
