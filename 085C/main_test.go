package main

import "testing"

func TestDama(t *testing.T) {
	n, y := 9, 45000
	wa, wb, wc := 4, 0, 5
	if ga, gb, gc := dama(n, y); ga != wa || gb != wb || gc != wc {
		t.Fatalf("want is %v %v %v , got is %v %v %v", wa, wb, wc, ga, gb, gc)
	}
}
