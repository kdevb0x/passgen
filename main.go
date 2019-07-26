package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	      	       c    to include capitols [A-Z]
	      	       l    to include lowercase [a-z]
	      	       n    to include numbers [0-9]
	      	       s    to include symbols [!@#$%&*]

Examples:
     passgen -l 10 -i ln
     passgen -i cls -o pass.txt -q
     passgen -l 24 -i lcns -v --output=/home/me/totallyNotAPassword.txt`

var (
	usage = usageString

	includeFlags []string // char groups to include in generation
	length       int      // number of character in generated string
	filepath     string   // path for -o flag
	forcef       bool     // force overwrite output file (if already exists)
	silent       bool
	verify       bool // TODO: consider removing this and make it mandatory
)

func getFlags() {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringArrayVarP(&includeFlags, "include", "i", []string{"l"}, "`character classes` to include in the generation")

	pflag.StringVarP(&filepath, "output", "o", "", "write output string to given FILE")
	pflag.BoolVarP(&forcef, "force", "f", false, "overwrite file instead of appending if it already exists (implies -o)")
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

	for i := 0; i < len(includeFlags); i++ {
		cons := strings.Split(includeFlags[i], "")
		for _, con := range cons {
			switch con {
			case "l":
				constraints["lower"] = true
			case "c":
				constraints["upper"] = true
			case "n":
				constraints["number"] = true
			case "s":
				constraints["symbol"] = true
			case " ":
				fallthrough
			default:
				fmt.Printf("%v: invalid argument", pflag.Args()[1:])
				fmt.Print(usage)
				os.Exit(1)
			}
		}
	}

}

// writeFile writes the generated string to the filepath indicated by path.
// If path == "", it defaults to `~/passgen_YEAR\MONTH\DAY.txt`
// NOTE: backslashes ('\') are for readability and are NOT in the actual filename.
func writeFile(pass string, path string) error {
	var fname string // final filename
	// throw error is more than 1 filepath
	if path == "" {
		year, month, day := time.Now().Date()
		dateval := strings.Join([]string{string(year), month.String(), string(day)}, "")
		path = fmt.Sprintf("passgen_%s", dateval)

		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		fname = homedir + path
	}

	var file *os.File
	var _, err = os.Stat(fname)
	if err != nil {

		// check err to see if file exists or not.
		// since we only care if the file exists or not,
		// if err is some other error, we just return it as is.
		switch {
		case os.IsNotExist(err):
			log.Printf("creating file %s\n", fname)
			file, err = os.Create(fname)
			if err != nil {
				return errors.New(string("failed to create file; error:" + err.Error()))
			}
			defer file.Close()
		case os.IsExist(err):
			if !forcef {
				fmt.Printf("%s already exists, appending to end of file. To overwrite instead, use -f, --force\n", fname)
				file, err = os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_APPEND, os.ModeAppend)
				if err != nil {
					return err
				}
				defer file.Close()

			}
		default:
			// we only care if the file exists or not, so if
			// err is some other error, we just return it.
			return err
		}
	}

	file, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0600)
	defer file.Close()

	var nline = pass + "\n"
	_, err = file.WriteString(nline)
	if err != nil {
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
	if len(filepath) > 0 {
		// NOTE: commenting this out for now so as not to duplicate
		// functionality, but it may be good for filepath
		// verification or similar, so I'll keep it around (for now).

		/*
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

		 		// log.Fatal(err)
		 	}
		 	os.Exit(0)
		 }
		*/

		if writeFile(pass, filepath); err != nil {
			log.Fatal(err)
		}

	}
	if err := writeFile(pass, ""); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
