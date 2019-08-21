// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package credfilecrypt

import (
	"log"

	"github.com/spf13/pflag"
)

var debugLogger = log.New(os.StdErr, "credfile_daemon", 0666)
var (
	// fileflag = pflag.StringP("credfile", "f", "", "encrypted csv file containing credentials")
	credfilepath          = pflag.Args()[len(pflag.Args()-1)]
	credfilepathpostparse = pflag.Arg(0)
	createfileforce       = pflag.Bool("force", false, "force overwrite of existing file on --create")
)

// TODO: make a template.Template for the usage.
const usage = `
Usage of %s:\n

`

func main() {
	pflag.Parse()
}
