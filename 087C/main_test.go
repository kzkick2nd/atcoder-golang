package main

import "testing"

func TestSeek(t *testing.T) {
	cases := map[string]struct {
		n    int
		a, b []int
		want int
	}{
		"1": {n: 5, a: []int{3, 2, 2, 4, 1}, b: []int{1, 2, 2, 2, 1}, want: 14},
		"2": {n: 4, a: []int{1, 1, 1, 1}, b: []int{1, 1, 1, 1}, want: 5},
		"3": {n: 7, a: []int{3, 3, 4, 5, 4, 5, 3}, b: []int{5, 3, 4, 4, 2, 3, 2}, want: 29},
		"4": {n: 1, a: []int{2}, b: []int{3}, want: 5},
	}

	for i, tc := range cases {
		tc := tc
		t.Run(i, func(t *testing.T) {
			if got := seek(tc.n, tc.a, tc.b); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
