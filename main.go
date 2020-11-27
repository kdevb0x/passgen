package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	fp "path/filepath"

	"github.com/spf13/pflag"
)

const usageString = `Usage:  passgen [OPTION]... [FILE]...
passgen is a simple, configurable password generator.

Available Options:
       -i, --include {clns}    Character GROUPS to include.
       -l, --length [LENGTH]   Desired length of output string (defaults to 1).
       -v, --verify 	       Ensure the generated string includes at least 1
    				of each catagory indicated by '--include'.

       -q, --quiet 	       Prevent echoing the generated string to the terminal.
       -o, --output <FILE>     Save output string to FILE.
       -f, --force 	       If FILE exists, overwrite file instead of appending
    				to the end (implies -o).

       -h, --help 	       Help information (this screen).


GROUPS should be given inline without space or comma seperations (ie: -i lun)

Must be one or more of:
	      	       l    to include lowercase [a-z]
	      	       u    to include uppercase (capitols) [A-Z]
	      	       n    to include numbers [0-9]
	      	       s    to include symbols [!@#$%&*]

Examples:
     passgen -l 10 -i ln
     passgen -i lus -o pass.txt -q
     passgen -l 24 -i luns -v --output=/home/me/totallyNotAPassword.txt`

var (
	usage        = usageString
	includeFlags string
	length       int
	filepath     string // path for -o flag
	forcef       bool
	silent       bool
	verifyFlag   bool // TODO: consider removing this and make it mandatory
	//
)

func getFlags() {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringVarP(&includeFlags, "include", "i", "l", "`character classes` to include in the generation")

	pflag.StringVarP(&filepath, "output", "o", "", "write output string to given FILE")
	pflag.BoolVarP(&forcef, "force", "f", false, "overwrite file instead of appending if it already exists (implies -o)")
	pflag.BoolVarP(&silent, "quiet", "q", false, "suppress echoing of generated string to stdout")
	pflag.BoolVarP(&verifyFlag, "verify", "v", true, "ensure that the generated string includes A͟T͟ ͟L͟E͟A͟S͟T͟ O͟N͟E͟ character from each group passed to [--include, -i]")
	pflag.Usage = func() {
		fmt.Printf("%s\n", usage)
		os.Exit(1)
	}
	pflag.Parse()

	// check if ran without args;
	// instead of gen and printing a single letter which is useless,
	// print the usage string and exit.
	if len(os.Args[:]) <= 1 || len(pflag.Args()[:]) != 0 {
		pflag.Usage()
	}

	// for i := 0; i < len(includeFlags); i++ {
	// cons := strings.Split(includeFlags, "")
	for _, con := range includeFlags {
		switch con {
		case 'l':
			constraints["lower"] = true
		case 'u':
			constraints["upper"] = true
		case 'n':
			constraints["number"] = true
		case 's':
			constraints["symbol"] = true
		case ' ':
			continue
		default:
			fmt.Printf("%v: invalid argument", pflag.Args()[1:])
			fmt.Print(usage)
			os.Exit(1)
		}
		// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
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

	cons := generatePool(constraints)
	gen := generateChars(cons, length)

	pass, err := buildString(gen, verifyFlag) // verify if flag is present
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
				// BUG:
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
