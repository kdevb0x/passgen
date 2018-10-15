package main // import "github.com/kidoda/passgen"

import (
	"flag"

	"github.com/codegangsta/cli"
	mow "github.com/jawher/mow.cli"
)

var constraint = make(map[string]bool)

/*
func getflags() {
        pflag.StringToStringVarP(&constraints, "include", "i", constr, `Elements to include in generation. Change to 'y' to include. Default options are: 'characters':'n', 'uppercase':'n', 'lowercase':'n', 'numbers':'n'`)

        pflag.Parse()
}

*/

func execAction(ctx *cli.Context) error {

}

func setupApp() *mow.Cli {
	app := cli.NewApp()
	app.Name = "passgen"
	app.Description = "A configurable constraint-based password string generator"
	app.Flags = []cli.Flag{cli.Flag{"L length", 8, "length of the generated password string"}}

	flag.Visit(func(f *flag.Flag) { constraint[f.Name] = true })
}
func main() {

}
