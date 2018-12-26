// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"testing"
)

func TestEncryptAES(t *testing.T) {

}

func TestHashBcrypt(t *testing.T) {
	var pw = NewPW(nil)
	pwhash := HashBcrypt(pw)
}

func TestNewPW(t *testing.T) {
	pwstring := []byte("testpw")
	hash := NewPW(pwstring)
}
