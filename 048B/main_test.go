package main

import "testing"

func TestBetween(t *testing.T) {
	cases := map[string]struct {
		a, b, x, want int
	}{
		"1": {a: 4, b: 8, x: 2, want: 3},
		"2": {a: 0, b: 5, x: 1, want: 6},
		"3": {a: 9, b: 9, x: 2, want: 0},
		"4": {a: 1, b: 1000000000000000000, x: 3, want: 333333333333333333},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := between(tc.a, tc.b, tc.x); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
