package main

import "testing"

func TestDream(t *testing.T) {
	cases := map[string]struct {
		s, want string
	}{
		"1": {s: "erasedream", want: "YES"},
		"2": {s: "dreameraser", want: "YES"},
		"3": {s: "dreamerer", want: "NO"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := dream(tc.s); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
