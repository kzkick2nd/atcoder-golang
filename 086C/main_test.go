package main

import "testing"

func TestTravel(t *testing.T) {
	cases := map[string]struct {
		n          int
		time, x, y []int
		want       string
	}{
		"1": {
			n:    2,
			time: []int{3, 6},
			x:    []int{1, 1},
			y:    []int{2, 1},
			want: "Yes",
		},
		"2": {
			n:    1,
			time: []int{2},
			x:    []int{100},
			y:    []int{100},
			want: "No",
		},
		"3": {
			n:    2,
			time: []int{5, 100},
			x:    []int{1, 1},
			y:    []int{1, 1},
			want: "No",
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := travel(tc.n, tc.time, tc.x, tc.y); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}

}
