package main

import "testing"

func TestOn(t *testing.T) {
	cases := map[string]struct {
		s      string
		expect int
	}{
		"1": {s: "oxo", expect: 900},
		"2": {s: "ooo", expect: 1000},
		"3": {s: "xxx", expect: 700},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := on(tc.s); actual != tc.expect {
				t.Fatalf("expect is %d, acutual is %d", tc.expect, actual)
			}
		})
	}
}
