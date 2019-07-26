// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"encoding/hex"
	"log"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	Username string
	Email    string // recovery email address
	email    *mail.Address
	Password string // base64 encoded hash
}

// AddPassword turns given password and optional salt into base64 encoded password
// hash, adding it to the credential.
// if multiple salts are provided, the are concatenated together and used as a single salt.
func (c *Credential) AddPassword(password []byte) error {
	phash, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return err
	}
	c.Password = hex.EncodeToString(phash)
	return nil
}

func NewCredential(username, email string) *Credential {
	var c = new(Credential)
	addr, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("error parsing address %s ERROR: %s, Using raw string\n", email, err.Error())
		c.Email = email
	}
	c.Username = username
	c.email = addr
	return c

}
