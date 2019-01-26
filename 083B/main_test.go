package main

import "testing"

func TestSum(t *testing.T) {
	cases := map[string]struct {
		n, a, b, expect int
	}{
		"1": {n: 20, a: 2, b: 5, expect: 84},
		"2": {n: 10, a: 1, b: 2, expect: 13},
		"3": {n: 100, a: 4, b: 16, expect: 4554},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := sum(tc.n, tc.a, tc.b); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
