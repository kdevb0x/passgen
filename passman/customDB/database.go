// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package customDB

import (
	"errors"
	"sync"
)

type status int

const (
	_ status = iota
	unlocked
	locked
	uninitialized
)

type EncryptedBuffer struct {
	mu   *sync.Mutex
	buff []byte
	stat status
}

// NewEncryptedBuffer creates and initializes a new EncryptedBuffer optionally
// using buf as its initial contents. It takes ownership of buf, and the caller
// should not use buf after this call.
func NewEncryptedBuffer(buf ...[]byte) (*EncryptedBuffer, error) {
	if len(buf) != 0 {
		if len(buf) > 1 {
			return nil, errors.New("buf must be a single []byte, or nil")
		}
		return &EncryptedBuffer{buff: buf[0]}, nil
	}
	return &EncryptedBuffer{buff: make([]byte, 1024)}, nil

}

// Write appends len(p) bytes from p into e's buffer, growing buffer if needed,
// then returns the new length of e.buff, and a nil error.
func (e *EncryptedBuffer) Write(p []byte) (n int, err error) {
	n = copy(e.buff, p)
	return
}

func (e *EncryptedBuffer) Read(p []byte) (n int, err error) {
	n = copy(p, e.buff[:])
	return n, nil
}

func (e *EncryptedBuffer) Bytes() []byte {
	return e.buff[:]
}

func (e *EncryptedBuffer) Reset() {
	e.buff = e.buff[:0]
}

func (e *EncryptedBuffer) Unlock(key string) error {
	if e.IsUnlocked {
		return errors.New("EncryptedBuffer already unlocked")
	}
	e.stat = unlocked
	return nil
}

func (e *EncryptedBuffer) IsUnlocked() bool {
	if e.stat == unlocked {
		return true
	}
	return false
}
