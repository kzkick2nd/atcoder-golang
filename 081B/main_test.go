package main

import "testing"

func TestChange(t *testing.T) {
	cases := map[string]struct {
		n    int
		list []int
		want int
	}{
		"1": {n: 3, list: []int{8, 12, 40}, want: 2},
		"2": {n: 4, list: []int{5, 6, 8, 10}, want: 0},
		"3": {n: 4, list: []int{382253568, 723152896, 37802240, 379425024, 404894720, 471526144}, want: 8},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := change(tc.n, tc.list); got != tc.want {
				t.Fatalf("want %v, got %v", tc.want, got)
			}
		})
	}
}
