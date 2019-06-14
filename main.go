package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Text to modify
	wordlist := flag.String("w", "", "wordlist")
	text := flag.String("t", "", "text: one,two,three,...")

	// Modifieres
	modifierNames := flag.String("m", "", "modifier names")
	modifiersFile := flag.String("mf", "", "modifiers file")

	// Extra
	info := flag.Bool("i", false, "print information")

	flag.Parse()

	var strs []string

	if *wordlist != "" {
		data, err := ioutil.ReadFile(*wordlist)

		if err != nil {
			log.Fatal("Filename not found")
			os.Exit(1)
		}

		strs = strings.Split(string(data), "\n")
	} else if *text != "" {
		strs = strings.Split(*text, ",")
	} else {
		panic("Wordlist or text must be specified")
	}

	if *modifierNames == "" && *modifiersFile == "" {
		panic("Modifiers must be specified")
	}

	instructions := *modifierNames

	if *modifiersFile != "" {
		data, err := ioutil.ReadFile(*modifiersFile)

		if err != nil {
			log.Fatal("Error opening modifiers file")
			os.Exit(1)
		}

		instructions = strings.ReplaceAll(string(data), "\n", ";")
		instructions = instructions[:len(instructions)-1] // remove last ;
		fmt.Println(instructions)
	}

	strs = removeEmptyStrings(strs)
	modifiers, err := parse(instructions)

	if err != nil {
		log.Fatal(err)
	}

	if *info {
		ls, lm := len(strs), len(modifiers)
		fmt.Printf("%d words will be modified.\n", ls)
		fmt.Printf("%d modifiers will be applied.\n", lm)
		fmt.Printf("%d words will be generated.\n", ls*lm)
		fmt.Printf("%d total words.\n", ls*lm+ls)
		os.Exit(0)
	}

	for _, str := range strs {
		fmt.Println(str)

		modifiedStrs := applyModifiers(modifiers, str)

		for _, m := range modifiedStrs {
			if m != "" {
				fmt.Println(m)
			}
		}
	}
}
