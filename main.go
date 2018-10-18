package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const usageString = `Usage:  passgen [OPTION]... [FILE]...
passgen is a simple, configurable password generator written in Go.

Available Options:
	-i, --include <TYPES>    character types to include
	-l, --length [LENGTH]    desired length of output string (defaults to 1)
	-o, --output <FILE>	 save output string to FILE

	-h, --help 		 help information (this screen)

TYPES should be given with no spaces or commas seperating them (ie: -i lun)
and must be one or more of:

	l 	to include lowercase (a-z)
	u	to include uppercase (A-Z)
	n 	to include numbers (0-9)
	s 	to include symbols (!@#$%&*)

Examples:
	passgen -l 10 -i ln
	passgen -i lus -o pass.txt
	passgen -l 24 -i luns --output=/home/me/totallyNotAPassword.txt`

var (
	usage        = usageString
	includeFlags []string
	length       int
)

func getFlags() *os.File {
	pflag.IntVarP(&length, "length", "l", 1, "`length` of the generated output string")
	pflag.StringArrayVarP(&includeFlags, "include", "i", nil, "character `classes` to include in the generation")

	filepath := pflag.StringP("output", "o", "", "write output string to given FILE")

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
				fmt.Printf("%v: invalid argument", os.Args[0])
				defer fmt.Print(usage)
				os.Exit(1)
			}
			// flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
		}
	}
	if *filepath != "" {
		file, err := os.Create(*filepath)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		return file
	}
	return nil
}

func main() {
	outfile := getFlags()

	cons := checkConstraints(constraints)
	gen := generateChars(cons)

	pass, err := buildString(gen)
	if err != nil {
		log.Fatal(err)
	}

	if outfile != nil {
		_, err := outfile.WriteString(pass)
		if err != nil {
			log.Fatal("error writing generated string to file!:", err)
		}
		defer outfile.Close()
		os.Exit(0)
	}
	fmt.Println(pass)
	os.Exit(0)
}
