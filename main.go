package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

/*
// custom flag for the character classes that define constraints.
type flagList []string

func (f *flagList) Set(value string) error {
	*f = strings.Split(value, ",")
	return nil
}

func (f *flagList) String() string {
	return fmt.Sprintf("%v", *f)
}
*/
const usageString = `Usage:  passgen [OPTION]... [FILE]...
passgen is a simple, configurable password generator written in Go.

Available Options:
	-i, --include <TYPES>    character types to include
	-l, --length [LENGTH]    desired length of output string
	-o, --output <FILE>	 save output string to FILE

	-h, --help 		 help information (this screen)

TYPES must be one or more of [l]owercase, [u]ppercase, [n]umbers, or [s]ymbols.
(meaning: a-z, A-Z, 0-9, )`

var (
	usage        = flag.PrintDefaults
	includeFlags []string
	length       int
)

func getFlags() {

	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")

	pflag.StringArrayVarP(&includeFlags, "include", "i", nil, "character `classes` to include in the generation")
	pflag.Parse()
	for i := 0; i < len(includeFlags); i++ {
		cons := strings.Split(includeFlags[i], "")
		for _, con := range cons {
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
	}
	return
}

func main() {
	getFlags()

	cons := checkConstraints(constraints)
	gen := generateChars(cons)

	pass, err := buildString(gen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pass)
	os.Exit(0)
}
