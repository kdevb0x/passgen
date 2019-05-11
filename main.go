package main

import (
	"errors"
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
    -f, --force 	     if FILE exists, overwrite file instead of appending to the end (implies -o)
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
	filepath     string // path for -o flag
	forcef       bool
	silent       bool
	verify       bool // TODO: consider removing this and make it mandatory
	//
)

func getFlags() {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringArrayVarP(&includeFlags, "include", "i", []string{"l"}, "`character classes` to include in the generation")

	pflag.StringVarP(&filepath, "output", "o", "", "write output string to given FILE")
	pflag.BoolVarP(&forcef, "force", "f", false, "overwrite file instead of appending if it already exists (implies -o)")
	pflag.BoolVarP(&silent, "quiet", "q", false, "suppress echoing of generated string to stdout")
	pflag.BoolVarP(&verify, "verify", "v", true, "ensure that the generated string includes A͟T͟ ͟L͟E͟A͟S͟T͟ O͟N͟E͟ character from each group passed to [--include, -i]")
	pflag.Usage = func() {
		fmt.Print(usage)
	}
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
			case " ":
				continue
			default:
				fmt.Printf("%v: invalid argument", pflag.Args()[1:])
				fmt.Print(usage)
				os.Exit(1)
			}
			// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
		}
	}

}
func writeFile(pass string) error {
	/*
		var outfile *os.File


		if filepath != "" {
			outfile, err := os.Create(filepath)
			if err != nil {
				return fmt.Errorf("Error creating output file:", err)
			}
			defer outfile.Close()

		}
		if outfile != nil {
			n, err := outfile.WriteString(pass)
			if err != nil || n < len(pass) {
				log.Fatal("error writing generated string to file!:", err)
			}
		}
	*/
	// if err := ioutil.WriteFile(filepath, []byte(pass), 0600); err != nil {
	var f *os.File
	f, err := os.Create(filepath)
	if err != nil {
		switch err {
		case os.ErrExist:
			fmt.Printf(`%s already exists, appending to end of file.
			To overwrite instead, use -f, --force flag`, filepath)

			f, err := os.Open(filepath)
			if err != nil {
				return err
			}
			defer f.Close()
		case os.ErrPermission:
			return errors.New("unable to write file, invalid permissions:" + err.Error())

		}
	}
	defer f.Close()
	n, err := f.WriteString(pass)
	if err != nil || n != len(pass) {
		return fmt.Errorf("unable to write to %s, reason: %s", filepath, err.Error())
	}
	return nil

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
	if filepath != "" {
		if err := writeFile(pass); err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}
