package main

import "testing"

func TestAlpha(t *testing.T) {
	cases := map[string]struct {
		s, expect string
	}{
		"1": {s: "atcoderregularcontest", expect: "b"},
		"2": {s: "abcdefghijklmnopqrstuvwxyz", expect: "None"},
		"3": {s: "fajsonlslfepbjtsaayxbymeskptcumtwrmkkinjxnnucagfrg", expect: "d"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := alpha(tc.s); actual != tc.expect {
				t.Fatalf("expect is %s, actual is %s", tc.expect, actual)
			}
		})
	}
}
