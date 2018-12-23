package main

import "testing"

func TestTrain(t *testing.T) {
	cases := map[string]struct {
		n, want int
	}{
		"1": {n: 3, want: 6},
		"2": {n: 10, want: 3628800},
		"3": {n: 100000, want: 457992974},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if got := train(tc.n); got != tc.want {
				t.Fatalf("want is %v, got is %v", tc.want, got)
			}
		})
	}
}
