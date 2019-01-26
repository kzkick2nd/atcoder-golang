package main

import "testing"

func TestHarshad(t *testing.T) {
	cases := map[string]struct {
		n      int
		expect string
	}{
		"1": {n: 12, expect: "Yes"},
		"2": {n: 57, expect: "No"},
		"3": {n: 148, expect: "No"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := harshad(tc.n); actual != tc.expect {
				t.Fatalf("expect is %s, actual %s", tc.expect, actual)
			}
		})
	}
}
