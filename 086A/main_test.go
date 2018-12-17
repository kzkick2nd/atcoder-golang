package main

import "testing"

func TestEvenodd(t *testing.T) {
	cases := map[string]struct {
		a, b int
		want string
	}{
		"1": {a: 3, b: 4, want: "Even"},
		"2": {a: 1, b: 21, want: "Odd"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := evenodd(tc.a * tc.b); got != tc.want {
				t.Fatalf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}
