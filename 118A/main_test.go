package main

import "testing"

func TestPlusMinus(t *testing.T) {
	cases := map[string]struct {
		a, b   int
		expect int
	}{
		"1": {a: 4, b: 12, expect: 16},
		"2": {a: 8, b: 20, expect: 12},
		"3": {a: 1, b: 1, expect: 2},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := pm(tc.a, tc.b); actual != tc.expect {
				t.Fatalf("expext is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
