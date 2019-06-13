package main

import "testing"

func TestSelectSimpleModifier(t *testing.T) {
	m, err := selectSimpleModifier("w")

	if m != nil && err == nil {
		t.Error("should return an error")
	}

	m, err = selectSimpleModifier("l")

	if m == nil && err != nil {
		t.Error("should return modifier")
	}

	r := m("HEllo")

	if r != "hello" {
		t.Error("selected modifier invalid")
	}
}

func TestSelectModifierWithArgument(t *testing.T) {
	m, err := selectModifierWithArgument("w", "123")

	if m != nil && err == nil {
		t.Error("should return an error")
	}

	m, err = selectModifierWithArgument("a", "123")

	if m == nil && err != nil {
		t.Error("should return modifier")
	}

	r := m("Hello")

	if r != "Hello123" {
		t.Error("selected modifier invalid")
	}
}

func TestSelectModifier(t *testing.T) {
	// Simple
	m, _ := selectModifier("u")
	r := m("Hello")

	if r != "HELLO" {
		t.Error("selected modifier invalid")
	}

	// With args
	m, _ = selectModifier("p:qwerty")
	r = m("Hello")

	if r != "qwertyHello" {
		t.Error("selected modifier invalid")
	}

	m, _ = selectModifier("p:12:ab")
	r = m("Hello")

	if r != "12:abHello" {
		t.Error("selected modifier invalid")
	}
}

func TestConcatModifiers(t *testing.T) {
	// Concat
	m, _ := selectModifier("<u,d>")
	r := m("Hello")

	if r != "HELLOHELLO" {
		t.Error("selected concatenated modifier invalid")
	}

	m, _ = selectModifier("<u,d>")
	r = m("Hello")

	if r != "HELLOHELLO" {
		t.Error("selected concatenated modifier invalid")
	}

	m, _ = selectModifier("<a:!,p:!,u>")
	r = m("Hello")

	if r != "!HELLO!" {
		t.Error("selected concatenated modifier invalid")
	}
}

func TestParseComplexModifiers(t *testing.T) {
	// Error
	modifiers, err := parse("a:8;u;w")

	if err == nil {
		t.Error("should return an error")
	}

	// Complex
	modifiers, _ = parse("l;u;a:!")
	str := "Hello"

	if modifiers[0](str) != "hello" {
		t.Error("error")
	}

	if modifiers[1](str) != "HELLO" {
		t.Error("error")
	}

	if modifiers[2](str) != "Hello!" {
		t.Error("error")
	}
}

func TestModifierInBothMaps(t *testing.T) {
	m, _ := selectModifier("d")
	r := m("Hello")

	if r != "HelloHello" {
		t.Error("error")
	}

	m, _ = selectModifier("d:-")
	r = m("Hello")

	if r != "Hello-Hello" {
		t.Error("error")
	}
}

func TestParsingError(t *testing.T) {
	_, err := selectModifier("l:123")
	if err == nil {
		t.Error("should return an error")
	}

	_, err = selectModifier("a")
	if err == nil {
		t.Error("should return an error")
	}

	_, err = selectModifier("<l,a>")
	if err == nil {
		t.Error("should return an error")
	}
}
