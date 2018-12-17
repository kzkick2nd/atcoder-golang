package main

import "testing"

func TestPattern(t *testing.T) {
	cases := map[string]struct {
		a, b, c, x int
		want       int
	}{
		"1": {a: 2, b: 2, c: 2, x: 100, want: 2},
		"2": {a: 5, b: 1, c: 0, x: 150, want: 0},
		"3": {a: 30, b: 40, c: 50, x: 6000, want: 213},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := coins(tc.a, tc.b, tc.c, tc.x); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
