// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	bd "github.com/go-bindata/go-bindata"
	"github.com/awnumar/memguard"
)

const defaultPassLen = 16
const defaultConstraints = "luns"


type fileHandler struct {
	// internal buffer for enclave io.
	f           *memguard.LockedBuffer
	// encrypted enclave of ciphertext.
	e memguard.Enclave
	// reports if f has data.
	hasData bool
	// reports if enclave is sealed.
	sealed bool
}

func NewFileHandlerFromFile(file string) *fileHandler {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("<simplePwFile> error reading existing file %s. %w\n", file, err)
	}
	h := &fileHandler{f: memguard.NewBuffer(len(f)), hasData: true}
	h.f.Move(f)
	h.f.Freeze()
	return h
}

func NewFileHandlerFromInternal() (*fileHandler, error) {
	f, err := pwTxt()
	if err != nil {
		return nil, err
	}
	h := &fileHandler{
		f: memguard.NewBufferFromBytes(f.bytes),
		hasData: true,
	}
	h.f.Melt()
	h.e = *memguard.NewEnclave(h.f.Bytes())

	return h, nil

}

type password struct {
	b []byte
	err string
}

func (p *password) Error() string {
	return p.err
}

func (h *fileHandler) Add(provider, email string, password []byte) error {
	if len(password) == 0 {
		return errors.New("must provide password string to add")
	/*

		err := h.NewPassword(provider, email)
		if err != nil {
			return err
		}
		password = p
	}
	*/
	if !h.f.Mutable() {
		h.f.Melt()
	}

	b := new(bytes.Buffer)
	if _, err := b.WriteString(provider); err != nil {
		return err
	}
	if _, err := b.WriteString(email); err != nil {
		return err
	}
	if _, err := b.Write(password); err != nil {
		return err
	}
	h.f.Copy(b.Bytes())
	h.f.Freeze()

	return nil
}

func (h *fileHandler) ReadCipherTxtFromBin() error {
	var _ bd.Asset
	// derive cwd and make temp file for enc/decryption
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error calculating working dir: %w\n", err)
	}

	tmp, err := ioutil.TempFile(cwd, ".tmp")
	if err != nil {
		return fmt.Errorf("unable to create temp file! error: %w\n", err)
	}



}

func main() {

}
/*
func (h *fileHandler) NewPassword(with string, email string) error {
	var p password
	s, err := h.genPassword(defaultPassLen, defaultConstraints)
}

func (h *fileHandler) genPassword(length int, constraints string) ([]byte, error) {
	passgen
}

*/
