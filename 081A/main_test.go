package main

import "testing"

func TestOne(t *testing.T) {
	cases := map[string]struct {
		in     string
		expect int
	}{
		"1": {in: "101", expect: 2},
		"2": {in: "000", expect: 0},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := one(tc.in); actual != tc.expect {
				t.Fatalf("expect is %d, actual is %d", tc.expect, actual)
			}
		})
	}
}
