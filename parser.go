package main

import (
	"fmt"
	"strings"
)

var simpleModifiers = map[string]Modifier{
	// Simple
	"l": lower,
	"u": upper,
	"c": capitalize,
	"d": duplicate,
}

var modifierFactories = map[string]ModifierFactory{
	// With args
	"a":  appendStr,
	"p":  prependStr,
	"d":  duplicateWithSeparator,
	"rp": removePrefix,
	"rs": removeSuffix,
}

// Parsers
func parse(config string) ([]Modifier, error) {
	instructions := strings.Split(config, ";")
	modifiers := make([]Modifier, 0, len(instructions))

	for _, i := range instructions {
		modifier, err := selectModifier(i)

		if err != nil {
			return nil, err
		}

		modifiers = append(modifiers, modifier)
	}

	return modifiers, nil
}

func parseConcat(c string) (Modifier, error) {
	instructions := strings.Split(c, ",")
	modifiers := make([]Modifier, 0, len(instructions))

	for _, i := range instructions {
		modifier, err := selectModifier(i)

		if err != nil {
			return nil, err
		}

		modifiers = append(modifiers, modifier)
	}

	return concat(modifiers), nil
}

// Selectors
func selectModifier(i string) (Modifier, error) {
	if strings.Index(i, "<") == 0 && strings.LastIndex(i, ">") == (len(i)-1) {
		c := i[1 : len(i)-1]
		return parseConcat(c)
	} else if strings.Contains(i, ":") {
		s := strings.SplitN(i, ":", 2)

		if len(s) != 2 {
			return nil, fmt.Errorf("Parsing error in modifier with argument")
		}

		return selectModifierWithArgument(s[0], s[1])
	} else {
		return selectSimpleModifier(i)
	}
}

func selectSimpleModifier(name string) (Modifier, error) {
	modifier, ok := simpleModifiers[name]

	if !ok {
		return nil, fmt.Errorf("Modifier %s does not exist", name)
	}

	return modifier, nil
}

func selectModifierWithArgument(name string, arg string) (Modifier, error) {
	modifierFactory, ok := modifierFactories[name]

	if !ok {
		return nil, fmt.Errorf("Modifier %s does not exist", name)
	}

	return modifierFactory(arg), nil
}
