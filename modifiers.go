package main

import (
	"strings"
)

type Modifier func(str string) string
type ModifierFactory func(arg string) Modifier

func concat(modifiers []Modifier) Modifier {
	return func(str string) string {
		for _, m := range modifiers {
			str = m(str)
		}
		return str
	}
}

func applyModifiers(modifiers []Modifier, str string) []string {
	strs := make([]string, 0)
	for _, modifier := range modifiers {
		m := modifier(str)

		if m != "" {
			strs = append(strs, m)
		}
	}
	return strs
}

// Simple modifiers
func upper(str string) string {
	return strings.ToUpper(str)
}

func lower(str string) string {
	return strings.ToLower(str)
}

func capitalize(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func duplicate(str string) string {
	return str + str
}

func invert(str string) string {
	r := make([]rune, len(str))

	for i, c := range str {
		r[len(str)-i-1] = c
	}
	return string(r)
}

// Modifiers with args
func appendStr(a string) Modifier {
	return func(str string) string {
		return str + a
	}
}

func prependStr(p string) Modifier {
	return func(str string) string {
		return p + str
	}
}

func duplicateWithSeparator(sep string) Modifier {
	return func(str string) string {
		return str + sep + str
	}
}

func removePrefix(prefix string) Modifier {
	return func(str string) string {
		return strings.TrimPrefix(str, prefix)
	}
}

func removeSuffix(suffix string) Modifier {
	return func(str string) string {
		return strings.TrimSuffix(str, suffix)
	}
}
