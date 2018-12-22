package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const usageString = `passgen is a simple, configurable password generator written in Go.

Usage:  passgen [OPTION]... [FILE]...


Available Options:
    -i, --include {clns}     character GROUPS to include
    -l, --length [LENGTH]    desired length of output string (defaults to 1)

    -v, --verify 	     ensure the generated string includes at least 1 of each catagory indicated by '--include'

    -q, --quiet 	     prevent echoing the generated string to the terminal
    -o, --output <FILE>	     save output string to FILE
    -h, --help 		     help information (this screen)


GROUPS should be given inline without space or comma seperations (ie: -i lun)

Must be one or more of:
	    c	    to include capitol (A-Z)
	    l 	    to include lowercase (a-z)
	    n 	    to include numbers (0-9)
	    s 	    to include symbols (!@#$%&*)

Examples:
	passgen -l 10 -i ln (Generates a 10 char string of characters from [a-z0-9].)
	passgen -i cls -o pass.txt -q  (Generates a single char from [A-Za-z!@#$%&*] and writes it to the file: 'pass.txt' without echoing it to terminal.)
	passgen -l 24 -i lcns -v --output=/home/me/totallyNotAPassword.txt  (Generates a 24 char string, verified to include at least 1 from each of: [a-z], [A-Z], [0-9], and [!@#$%&*], writes it to the file: '~/totallyNotAPassword.txt', creating file if necessary, then echoes it to the terminal.)`

var (
	usage        = usageString
	includeFlags []string
	length       int
	filepath     string
	silent       bool
	verify       bool
)

func getFlags() {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringArrayVarP(&includeFlags, "include", "i", []string{"l"}, "`character classes` to include in the generation")

	pflag.StringVarP(&filepath, "output", "o", "", "write output string to given FILE")
	pflag.BoolVarP(&silent, "quiet", "q", false, "suppress echoing of generated string to stdout")
	pflag.BoolVarP(&verify, "verify", "v", false, "ensure that the generated string includes A͟T͟ ͟L͟E͟A͟S͟T͟ O͟N͟E͟ character from each group passed to [--include, -i]")
	pflag.Usage = func() {
		fmt.Print(usage)
	}
	pflag.Parse()

	for i := 0; i < len(includeFlags)-1; i++ {
		cons := strings.Split(includeFlags[i], "")
		for _, con := range cons {
			switch con {
			case "l":
				constraints["lower"] = true
			case "c":
				constraints["capitol"] = true
			case "n":
				constraints["number"] = true
			case "s":
				constraints["symbol"] = true
			default:
				fmt.Printf("%v: invalid argument", pflag.Args()[1:])
				fmt.Print(usage)
				os.Exit(1)
			}
			// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
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

	if filepath != "" {
		outfile, err := os.Create(filepath)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer outfile.Close()
		n, err := outfile.WriteString(pass)
		if err != nil || n < len(pass) {
			log.Fatal("error writing generated string to file!:", err)
		}
	}
	if silent == false {
		fmt.Println(pass)
	}

	os.Exit(0)
}
