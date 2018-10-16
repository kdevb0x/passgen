package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// custom flag for the character classes that define constraints.
type flagList []string

func (f *flagList) Set(value string) error {
	*f = strings.Split(value, ",")
	return nil
}

func (f *flagList) String() string {
	return fmt.Sprintf("%v", *f)
}

var (
	usage        = flag.PrintDefaults
	includeFlags flagList
	length       int
)

func getFlags() {
	flag.IntVar(&length, "L length", 1, "`length` of the generated output string")

	var include flagList
	flag.Var(&includeFlags, "I include", "character `classes` to include in the generation")
	flag.Parse()

	// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })

	return
}

func main() {
	getFlags()

	cons := checkConstraints(constraints)
	if len(constraints) <= 0 {
		log.Fatal("no constraints given!")
	}

	pass, err := generateChars(cons)
	if err != nil {
		log.Fatal(err)
	}
}
