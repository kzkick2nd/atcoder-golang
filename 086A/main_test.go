package main

import "testing"

func TestEvenodd(t *testing.T) {
	cases := map[string]struct {
		a, b   int
		expect string
	}{
		"1": {a: 3, b: 4, expect: "Even"},
		"2": {a: 1, b: 21, expect: "Odd"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := evenodd(tc.a * tc.b); actual != tc.expect {
				t.Fatalf("expect = %v, actual = %v", tc.expect, actual)
			}
		})
	}
}
