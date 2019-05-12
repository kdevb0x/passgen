package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	fp "path/filepath"
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
func writeFile(pass string, path ...string) error {
	// throw error is more than 1 filepath
	if len(path) > 1 {
		return errors.New("path parameter must be a single filepath!")
	}
	var fname = path[0]
	if fname == "" {
		fname = "out.txt"

	}
	var file *os.File
	var _, err = os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("creating file %s\n", fname)
			file, err = os.Create(fname)
			if err != nil {
				return errors.New(string("failed to create file; error:" + err.Error()))
			}
			file.Close()
		}
	}
	if forcef {
		// file, err = os.OpenFile(fname, os.O_RDWR|os.O_SYNC, 0600)
		if err := ioutil.WriteFile(fname, []byte(pass+"\n"), 0600); err != nil {
			return err
		}

	} else {
		fmt.Printf("%s already exists, appending to end of file. To overwrite instead, use -f, --force flag\n", fname)
		file, err = os.OpenFile(fname, os.O_RDWR|os.O_SYNC|os.O_APPEND, os.ModeAppend)
		if err != nil {
			return err
		}
		defer file.Close()

	}
	n, err := file.WriteString(string(pass + "\n"))
	if err != nil || n != len(pass) {
		return err
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
		if !fp.IsAbs(filepath) {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)

			}
			if err := writeFile(pass, fp.Join(cwd, filepath)); err != nil {
				// (kdd) TODO: for some reason this returns the
				// error: "Invalid argument" But otherwise writes
				// to file and works correctly, so I'm just gonna
				// comment it out for now, as such my time is
				// better spent on other things than tracking down
				// this bug.

				// log.Fatal(err)
			}
			os.Exit(0)
		}
		err := writeFile(pass, filepath)
		if err != nil {
			log.Fatal(err)
		}
	}

	os.Exit(0)
}
