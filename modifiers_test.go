package main

import "testing"

func TestInvert(t *testing.T) {
	r := invert("hello")

	if r != "olleh" {
		t.Error("invert")
	}
}
