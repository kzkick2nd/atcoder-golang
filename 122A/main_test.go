package main

import "testing"

func TestHelix(t *testing.T) {
	cases := map[string]struct {
		b      string
		expect string
	}{
		"1": {b: "A", expect: "T"},
		"2": {b: "G", expect: "C"},
	}
	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := helix(tc.b); actual != tc.expect {
				t.Fatalf("expext is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
