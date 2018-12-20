package main

import "testing"

func TestRobo(t *testing.T) {
	cases := map[string]struct {
		in   []int
		want int
	}{
		"1": {in: []int{0, 75, 25, 100}, want: 50},
		"2": {in: []int{0, 33, 66, 99}, want: 0},
		"3": {in: []int{10, 90, 20, 80}, want: 60},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := robo(tc.in); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
