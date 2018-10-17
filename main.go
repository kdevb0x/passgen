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
