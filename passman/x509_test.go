// Copyright 2018 kdevb0x Ltd. All rights reserved.
// This software is governed by the CeCILL license under French law
// and abiding by the rules of distribution of free software.
// The full license text can be found in the LICENSE file.

package passman

import (
	"testing"
)

func TestGetCertFromUrl(t *testing.T) {
	const gmailLogin = `https://accounts.google.com/signin/v2/identifier?service=mail&passive=true&rm=false&continue=https%3A%2F%2Fmail.google.com%2Fmail%2F&ss=1&scc=1&ltmpl=default&ltmplcache=2&emr=1&osid=1&flowName=GlifWebSignIn&flowEntry=ServiceLogin`

	cert, err := GetCertFromURL(gmailLogin)
	if err != nil {
		t.Fail()
	}

}
