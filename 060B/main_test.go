package main

import "testing"

func TestChoose(t *testing.T) {
	cases := map[string]struct {
		a, b, c int
		want    string
	}{
		"1": {a: 7, b: 5, c: 1, want: "YES"},
		"2": {a: 2, b: 2, c: 1, want: "NO"},
		"3": {a: 1, b: 100, c: 97, want: "YES"},
		"4": {a: 40, b: 98, c: 58, want: "YES"},
		"5": {a: 77, b: 42, c: 36, want: "NO"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := choose(tc.a, tc.b, tc.c); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
