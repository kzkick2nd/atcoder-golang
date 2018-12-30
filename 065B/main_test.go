package main

import "testing"

func TestTrain(t *testing.T) {
	cases := map[string]struct {
		n    int
		a    []int
		want int
	}{
		"1": {n: 3, a: []int{3, 1, 2}, want: 2},
		"2": {n: 4, a: []int{3, 4, 1, 2}, want: -1},
		"3": {n: 5, a: []int{3, 3, 4, 2, 4}, want: 3},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := train(tc.n, tc.a); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
