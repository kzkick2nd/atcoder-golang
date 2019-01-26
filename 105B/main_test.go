package main

import "testing"

func TestDonuts(t *testing.T) {
	cases := map[string]struct {
		n      int
		expect string
	}{
		"1": {n: 11, expect: "Yes"},
		"2": {n: 40, expect: "Yes"},
		"3": {n: 3, expect: "No"},
	}

	for i, tc := range cases {
		tc := tc
		t.Run(i, func(t *testing.T) {
			if actual := donuts(tc.n); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v.", tc.expect, actual)
			}
		})
	}
}
