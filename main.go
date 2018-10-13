package main // import "github.com/kidoda/passgen"

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var constraints = make(map[string]string, 10)

var constr = map[string]string{"characters": "n", "uppercase": "n", "lowercase": "n", "numbers": "n"}

func getflags() {
	pflag.StringToStringVarP(&constraints, "include", "i", constr, `Elements to include in generation. Change to 'y' to include. Default options are: 'characters':'n', 'uppercase':'n', 'lowercase':'n', 'numbers':'n'`)

	pflag.Parse()
}
