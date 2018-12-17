package main

import "testing"

func TestGame(t *testing.T) {
	cases := map[string]struct {
		n    int
		a    []int
		want int
	}{
		"1": {n: 2, a: []int{3, 1}, want: 2},
		"2": {n: 3, a: []int{2, 7, 4}, want: 5},
		"3": {n: 4, a: []int{20, 18, 2, 18}, want: 18},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := game(tc.n, tc.a); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
