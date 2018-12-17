package main

import "testing"

func TestSum(t *testing.T) {
	cases := map[string]struct {
		n, a, b, want int
	}{
		"1": {n: 20, a: 2, b: 5, want: 84},
		"2": {n: 10, a: 1, b: 2, want: 13},
		"3": {n: 100, a: 4, b: 16, want: 4554},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := sum(tc.n, tc.a, tc.b); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
