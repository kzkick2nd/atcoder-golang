package main

import "testing"

func TestOne(t *testing.T) {
	cases := map[string]struct {
		in   string
		want int
	}{
		"1": {in: "101", want: 2},
		"2": {in: "000", want: 0},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := one(tc.in); got != tc.want {
				t.Fatalf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}
