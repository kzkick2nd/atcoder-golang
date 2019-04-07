package main

import "testing"

func TestDiv(t *testing.T) {
	cases := map[string]struct {
		a, b, k  int
		expected int
	}{
		"1": {a: 8, b: 12, k: 2, expected: 2},
		"2": {a: 100, b: 50, k: 4, expected: 5},
		"3": {a: 1, b: 1, k: 1, expected: 1},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := divisor(tc.a, tc.b, tc.k); tc.expected != actual {
				t.Fatalf("expect is %v, actual is %v", tc.expected, actual)
			}
		})
	}
}
