package main

import "testing"

func TestPattern(t *testing.T) {
	cases := map[string]struct {
		a, b, c, x int
		expect     int
	}{
		"1": {a: 2, b: 2, c: 2, x: 100, expect: 2},
		"2": {a: 5, b: 1, c: 0, x: 150, expect: 0},
		"3": {a: 30, b: 40, c: 50, x: 6000, expect: 213},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := coins(tc.a, tc.b, tc.c, tc.x); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
