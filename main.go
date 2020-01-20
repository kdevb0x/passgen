package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
)

const usageString = `Usage:  passgen [OPTION]...
passgen is a simple, configurable password generator.

Available Options:
       -i, --include {luns}    Character GROUPS to include.
       -l, --length [LENGTH]   Desired length of output string (defaults to 1).
       -v, --verify 	       Ensure the generated string includes at least 1
    				of each catagory indicated by '--include'.

       -q, --quiet 	       Prevent echoing the generated string to the terminal.

       -h, --help 	       Help information (this screen).


GROUPS should be given inline without space or comma seperations (ie: -i luns)

Must be one or more of:
	      	       c    to include uppercase [A-Z]
	      	       l    to include lowercase [a-z]
	      	       n    to include numbers [0-9]
	      	       s    to include symbols [!@#$%&*]

Examples:
     passgen -l 10 -i ln
     passgen -i uls -o pass.txt -q
     passgen -l 24 -i luns -v `

var (
	usage = usageString

	includeFlags string // char groups to include in generation
	length       int    // number of character in generated string
	silent       bool
	verify       bool // TODO: consider removing this and make it mandatory
)

func getFlags() {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringVarP(&includeFlags, "include", "i", "l", "`character classes` to include in the generation")

	pflag.BoolVarP(&silent, "quiet", "q", false, "suppress echoing of generated string to stdout")
	pflag.BoolVarP(&verify, "verify", "v", true, "ensure that the generated string includes A͟T͟ ͟L͟E͟A͟S͟T͟ O͟N͟E͟ character from each group passed to [--include, -i]")
	pflag.Usage = func() {
		fmt.Printf("%s\n", usage)
		os.Exit(1)
	}
	pflag.Parse()

	// check if ran without args,
	// if so, print the usage string and exit.
	if len(os.Args[:]) <= 1 {
		pflag.Usage()
	}

	// for i := 0; i < len(includeFlags); i++ {
	// cons := strings.Split(includeFlags[i], "")
	for _, con := range includeFlags {
		switch con {
		case rune('l'):
			constraints["lower"] = true
		case rune('u'):
			constraints["upper"] = true
		case rune('n'):
			constraints["number"] = true
		case rune('s'):
			constraints["symbol"] = true
		case ' ':
			fallthrough
		default:
			fmt.Printf("error: found invalid argument(s)\n \r")
			fmt.Print(usage + "\n")
			os.Exit(1)
		}
	}

}

func main() {

	getFlags()

	cons := checkConstraints(constraints)
	gen := generateChars(cons)

	pass, err := buildString(gen)
	if err != nil {
		log.Fatal(err)
	}

	if !silent {
		fmt.Println(pass)
	}
	os.Exit(0)
}
