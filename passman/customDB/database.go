// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package customDB

import "errors"

// DB is the databse api

type DB interface {
	// TODO: describe the interface spec and guarentees
}

type encryptedBuffer struct {
	buff []byte
}

// NewEncryptedBuffer creates and initializes a new encryptedBuffer optionally
// using buf as its initial contents. It takes ownership of buf, and the caller
// should not use buf after this call.
func NewEncryptedBuffer(buf ...[]byte) (*encryptedBuffer, error) {
	if len(buf) != 0 {
		if len(buf) > 1 {
			return nil, errors.New("buf must be a single []byte, or nil")
		}
		return &encryptedBuffer{buff: buf[0]}, nil
	}
	return &encryptedBuffer{buff: make([]byte, 1024)}, nil

}

// Read implements the io.Reader interface. It copies len(p) bytes
// into the buffer and returns the number of bytes written and nil error.
func (e *encryptedBuffer) Read(p []byte) (n int, err error) {
	n = copy(e.buff, p)
	return
}

func (e *encryptedBuffer) Bytes()
