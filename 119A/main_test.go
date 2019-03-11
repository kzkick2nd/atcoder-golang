package main

import "testing"

func TestHeisei(t *testing.T) {
	cases := map[string]struct {
		s      string
		expect string
	}{
		"1": {s: "2019/04/30", expect: "Heisei"},
		"2": {s: "2019/11/01", expect: "TBD"},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			tc := tc
			if actual := heisei(tc.s); actual != tc.expect {
				t.Fatalf("expect is %v, actual is %v", tc.expect, actual)
			}
		})
	}
}
