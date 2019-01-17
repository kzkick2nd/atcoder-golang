package main

import "testing"

func TestChange(t *testing.T) {
	cases := map[string]struct {
		n      int
		a      []int
		expect int
	}{
		"1": {n: 3, a: []int{8, 12, 40}, expect: 2},
		"2": {n: 4, a: []int{5, 6, 8, 10}, expect: 0},
		"3": {n: 4, a: []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, expect: 8},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := change(tc.a); actual != tc.expect {
				t.Fatalf("expect is %d, actual is %d", tc.expect, actual)
			}
		})
	}
}
