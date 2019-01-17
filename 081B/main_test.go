package main

import "testing"

func TestChange(t *testing.T) {
	cases := map[string]struct {
		n      int
		list   []int
		expect int
	}{
		"1": {n: 3, list: []int{8, 12, 40}, expect: 2},
		"2": {n: 4, list: []int{5, 6, 8, 10}, expect: 0},
		"3": {n: 4, list: []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, expect: 8},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := change(tc.n, tc.list); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
