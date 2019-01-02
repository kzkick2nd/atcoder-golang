package main

import "testing"

func TestFace(t *testing.T) {
	cases := map[string]struct {
		n    int
		s    string
		want int
	}{
		"1": {n: 5, s: "WEEWW", want: 1},
		"2": {n: 12, s: "WEWEWEEEWWWE", want: 4},
		"3": {n: 8, s: "WWWWWEEE", want: 3},
	}

	for i, tc := range cases {
		tc := tc
		t.Run(i, func(t *testing.T) {
			if got := face(tc.n, tc.s); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
