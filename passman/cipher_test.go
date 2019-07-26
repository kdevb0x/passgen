// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const (
	tname  = "testUsername"
	temail = "test@email.com"
)

func TestCredential(t *testing.T) {
	cred := NewCredential(tname, temail)
	if err := cred.AddPassword([]byte("password")); err != nil {
		t.Fail()
	}
	if err := bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte("password")); err != nil {
		t.Fail()
	}

}
