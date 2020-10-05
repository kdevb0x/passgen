// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"crypto/rand"
	"crypto/aes"
	"io/ioutil"
	
	mg "github.com/awnumar/memguard"

	pgen "github.com/kdevb0x/passgen"
)

type cred struct {
	// top-level domain (ie: if login is login.google.com tld == "google.com")
	tld string
	email string 
	password Pass
}

type node interface {
	childL() node
	childR() node
	parent() node
}

// Passtree is a standard binary tree.
type passtree struct {
	parent node
	left node
	right node
	hash []byte 
}

func (pt *passtree) openDBFile(f string) (*mg.LockedBuffer, error) {
	db, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	b, err := mg.NewBufferFromEntireReader(db)
	if err != nil {
		return nil, err
	}
	
}
func (pt *passtree) parent() node {
	if pt.parent != nil {
		return pt.parent
	}

	// root node
	return nil
}

func (pt *passtree) childL() node {
	if pt.left != nil {
		return pt.left
	}

	return nil
}

func (pt *passtree) childR() node {
	if pt.right != nil {
		return pt.right
	}
	
	return nil
}