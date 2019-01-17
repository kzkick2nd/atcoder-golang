package main

import "testing"

func TestTGB(t *testing.T) {
	cases := map[string]struct {
		rgb    []int
		expect string
	}{
		"1": {rgb: []int{4, 3, 2}, expect: "YES"},
		"2": {rgb: []int{2, 3, 4}, expect: "NO"},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if actual := rgb(tc.rgb); actual != tc.expect {
				t.Fatalf("expect is %v. actual is %v", tc.expect, actual)
			}
		})
	}
}
