package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filename := flag.String("f", "", "wordlist")
	text := flag.String("t", "", "text: one,two,three,...")
	modifierNames := flag.String("m", "", "modifier names")
	flag.Parse()

	var strs []string

	if *filename != "" {
		data, err := ioutil.ReadFile(*filename)

		if err != nil {
			log.Fatal("Filename not found")
		}

		strs = strings.Split(string(data), "\n")
	} else if *text != "" {
		strs = strings.Split(*text, ",")
	} else {
		panic("Filename or text must be specified")
	}

	if *modifierNames == "" {
		panic("Modifiers must be specified")
	}

	strs = removeEmptyStrings(strs)

	modifiers, err := parse(*modifierNames)

	if err != nil {
		log.Fatal(err)
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
