package main

import "testing"

func TestFavSound(t *testing.T) {
	cases := map[string]struct {
		a, b, c int
		expect  int
	}{
		"1": {a: 2, b: 11, c: 4, expect: 4},
		"2": {a: 3, b: 9, c: 5, expect: 3},
		"3": {a: 100, b: 1, c: 10, expect: 0},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := favSound(tc.a, tc.b, tc.c); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
