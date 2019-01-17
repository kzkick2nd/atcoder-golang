package main

import "testing"

func TestNum(t *testing.T) {
	cases := map[string]struct {
		n, expect int
	}{
		"1": {n: 7, expect: 4},
		"2": {n: 32, expect: 32},
		"3": {n: 1, expect: 1},
		"4": {n: 100, expect: 64},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := num(tc.n); actual != tc.expect {
				t.Fatalf("expect is %d, actual is %d", tc.expect, actual)
			}
		})
	}
}
