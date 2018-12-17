package main

import (
	"reflect"
	"testing"
)

func TestDama(t *testing.T) {
	cases := map[string]struct {
		n, y int
		want []int
	}{
		"1": {
			n:    9,
			y:    45000,
			want: []int{0, 9, 0},
		},
		"2": {
			n:    20,
			y:    196000,
			want: []int{-1, -1, -1},
		},
		"3": {
			n:    1000,
			y:    1234000,
			want: []int{2, 54, 944},
		},
		"4": {
			n:    2000,
			y:    20000000,
			want: []int{2000, 0, 0},
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			a, b, c := dama(tc.n, tc.y)
			got := []int{a, b, c}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want is %v , got is %v", tc.want, got)
			}
		})
	}
}
