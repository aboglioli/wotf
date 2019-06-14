package main

import "testing"

func TestInvert(t *testing.T) {
	r := invert("hello")

	if r != "olleh" {
		t.Error("invert")
	}
}

func TestConditionalPrefixAndSuffix(t *testing.T) {
	// Prefix
	p := removePrefix("Qwe-")("Qwe-Hello")

	if p == "" || p != "Hello" {
		t.Error("prefix should be removed")
	}

	p = removePrefix("Qwe-")("Hello")

	if p != "" {
		t.Error("should return empty string")
	}

	// Suffix
	s := removeSuffix("-123")("Hello-123")

	if s == "" || s != "Hello" {
		t.Error("suffix should be removed")
	}

	s = removePrefix("-123")("Hello")

	if s != "" {
		t.Error("should return empty string")
	}
}
