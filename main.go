package main // import "github.com/kidoda/passgen"

import (
	"flag"
	"os"

	"github.com/codegangsta/cli"
)

// type flagList []string
type flagList struct {
	_ []string
}

var incFlagList flagList

var constraint = make(map[string]bool)

func (f *flagList) Set(value string) error {
	*f = strings.Split(value, ",")
	return nil
}

func (f *flagList) String() string {
	return fmt.Sprintf("%v", *f)
}

/*
func getflags() {
        pflag.StringToStringVarP(&constraints, "include", "i", constr, `Elements to include in generation. Change to 'y' to include. Default options are: 'characters':'n', 'uppercase':'n', 'lowercase':'n', 'numbers':'n'`)

        pflag.Parse()
}

*/

func execAction(ctx *cli.Context) error {

}

func getFlags() {

	var include = flag.Var(&incFlagList, "I include", "character `classes` to include in the generation")
	var strlenFlag = flag.Int("L length", 1, "desired length of the generated password string")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })

	return
}

func main() {
	getFlags()
	charclassLen := len(os.Args[2:])

	switch {
	case charclassLen <= 0:
		fmt.Printf("warning no constraints: can't generate a password string from nothing!")

	case charclassLen > 4:
		fmt.Printf("too many character classes given! Must be one of 'l,u,n,s' gave: %s", incFlagList)
	default:
		gen := NewPassGen(strlenFlag)
		gen.Generate()
	}
}
