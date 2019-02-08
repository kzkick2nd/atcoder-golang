package main

import "testing"

func TestLeap(t *testing.T) {
	cases := map[string]struct {
		t, x   float64
		expect float64
	}{
		"1": {t: 8, x: 3, expect: 2.6666666667},
		"2": {t: 99, x: 1, expect: 99.0000000000},
		"3": {t: 1, x: 100, expect: 0.0100000000},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := leap(tc.t, tc.x); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
