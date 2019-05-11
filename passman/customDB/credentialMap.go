// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package customDB

import (
	"bytes"
	"encoding/binary"
	"sync"
	"time"

	uuid "github.com/google/uuid" // or "github.com/nu7hatch/gouuid"
)

// LoginMetadata holds all the info that ties a password to an account/service
type LoginMetadata struct {
	UUID         uuid.UUID // used as key in credMap
	ValidUntil   *time.Time
	Username     string
	PasswordHash []byte // crypto/bcrypt hash of pw string
}

func (m *LoginMetadata) Expired() bool {
	now := time.Now()
	if !now.Before(*m.ValidUntil) {
		return false
	}
	return true
}

func (m *LoginMetadata) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.BigEndian, m)
	if err != nil {
		return
	}
	data = buf.Bytes()
	return

}

type Statuser interface {
	Status() status
}

// CredMap is a concurrent map of login credentials keyed by uuid.
type CredMap struct {
	mu    *sync.Mutex
	Creds map[string]LoginMetadata // keys are uuid strings
}

type Credential struct {
	buff      *EncryptedBuffer
	timestamp string
	metadata  *LoginMetadata
}

func (c *Credential) Status() status {
	return uninitialized
}

func NewCredential(ebuff *EncryptedBuffer) *Credential {

}
