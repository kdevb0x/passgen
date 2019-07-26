// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import "fmt"

type urlParseError struct {
	url string // url that was choked on
	err error  // the wrapped error
}

func (u *urlParseError) Error() string {
	return fmt.Sprintf("error encountered attempting to pars URL %s, ERROR: %s\n", u.url, u.err)
}
