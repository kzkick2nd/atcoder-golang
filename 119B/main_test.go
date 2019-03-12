package main

import "testing"

func TestYen(t *testing.T) {
	cases := map[string]struct {
		p      []pair
		expect float64
	}{
		"1": {p: []pair{{x: 10000, y: "JPY"}, {x: 0.10000000, y: "BTC"}}, expect: 48000.0},
		"2": {p: []pair{{x: 100000000, y: "JPY"}, {x: 100.00000000, y: "BTC"}, {x: 0.00000001, y: "BTC"}}, expect: 138000000.0038},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := yen(tc.p); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
