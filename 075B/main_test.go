package main

import (
	"reflect"
	"testing"
)

func TestMine(t *testing.T) {
	h, w := 3, 5
	maps := []string{".....", ".#.#.", "....."}
	want := []string{"11211", "1#2#1", "11211"}
	if got := mine(h, w, maps); !reflect.DeepEqual(got, want) {
		t.Fatalf("want is %v, got is %v", want, got)
	}
}
