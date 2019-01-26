package main

import "testing"

func TestSunuke(t *testing.T) {
	cases := map[string]struct {
		n, k   int
		l      []int
		expect int
	}{
		"1": {n: 5, k: 3, l: []int{1, 2, 3, 4, 5}, expect: 12},
		"2": {n: 15, k: 14, l: []int{50, 26, 27, 21, 41, 7, 42, 35, 7, 5, 5, 36, 39, 1, 45}, expect: 386},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := sunuke(tc.k, tc.l); actual != tc.expect {
				t.Fatalf("expect is %d, actual is %d", tc.expect, actual)
			}
		})
	}
}
