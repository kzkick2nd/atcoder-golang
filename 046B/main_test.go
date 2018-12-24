package main

import "testing"

func TestPaintBall(t *testing.T) {
	cases := map[string]struct {
		n, k int
		want int
	}{
		"1": {n: 2, k: 2, want: 2},
		"2": {n: 1, k: 10, want: 10},
		"3": {n: 10, k: 8, want: 322828856},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := paintBall(tc.n, tc.k); tc.want != got {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
