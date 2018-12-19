package main

import "testing"

func TestMochi(t *testing.T) {
	cases := map[string]struct {
		n    int
		d    []int
		want int
	}{
		"1": {n: 4, d: []int{10, 8, 8, 6}, want: 3},
		"2": {n: 3, d: []int{15, 15, 15}, want: 1},
		"3": {n: 7, d: []int{50, 30, 50, 100, 50, 80, 30}, want: 4},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := bucket(tc.d); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
