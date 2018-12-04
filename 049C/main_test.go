package main

import "testing"

func TestDream(t *testing.T) {
	s := "dreameraser"
	want := "YES"
	if got := dp(s); got != want {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
