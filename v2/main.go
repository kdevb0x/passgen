package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	flag.Var(&includeFlags, "I include", "character `classes` to include in the generation")
	flag.Parse()

	chars := strings.Split(includeFlags.String(), "")
	for _, con := range chars {
		switch con {
		case "l":
			constraints["lower"] = true
		case "u":
			constraints["upper"] = true
		case "n":
			constraints["number"] = true
		case "s":
			constraints["symbol"] = true
		default:
			fmt.Printf("include must be one or more of: {l|u|n|s} gave: %s", []string(includeFlags))
			flag.Usage()
			os.Exit(1)
		}
		// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
	}
	return
}

func main() {
	getFlags()

	cons := checkConstraints(constraints)
	pass, err := buildString(generateChars(cons))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pass)
	os.Exit(0)
}
