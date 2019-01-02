package main

import "testing"

func TestOp(t *testing.T) {
	cases := map[string]struct {
		in   string
		want string
	}{
		"1": {in: "1222", want: "1+2+2+2=7"},
		"2": {in: "0290", want: "0-2+9+0=7"},
		"3": {in: "3242", want: "3+2+4-2=7"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := op(tc.in); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
