// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package main

import (
	"encoding/binary"
	"os"
	"strings"

	"github.com/alexmullins/zip"
)

const DefaultRandPasswordLen = 32

type PwArchive struct {
	// File path on filesystem
	Fpath string

	// The hash of the password to compare against for decryption.
	pwhash []byte

	entries []AccountDetail

	// Whether the underlying encrypted zip has been decrypted.
	Decrypted bool
}

func CreateNewPwArchive(fpath string) (*PwArchive, error) {

	f, err := os.Create(fpath)
	if err != nil {
		return nil, err
	}
	var a = new(PwArchive)

	w := zip.NewWriter(f)

	a.Writer = w

	return a, nil

}

func (a *PwArchive) AddNewEntry(details AccountDetail) error {
	a.entries = append(a.entries, details)
	return nil
	//
	// Making AccountDetail strings and writing to temp files seems
	// uneeded: So lets just collect each pw as as an AccountDetail and
	// encode to binary and use that instead of strings.
	//

	/*
		var d []byte
		buff := bytes.NewBuffer(d)


		_, err := buff.WriteString(details.Provider + "\n")
		if err != nil {
			return err
		}
		_, err = buff.WriteString(details.Name + "\n")
		if err != nil {
			return err
		}
		_, err = buff.WriteString(string(details.Password) + "\n")
		if err != nil {
			return err
		}
		// f, err := ioutil.TempFile(os.TempDir(), "*")
		fp := os.TempDir() + "/" + details.Provider + ".pw"
		err = ioutil.WriteFile(fp, buff.Bytes(), os.ModeAppend)
		if err != nil {
			return err
		}

		return a.AddPwFile(fp, details.Password)
	*/

}

func (a *PwArchive) SaveAndClose(savepath string, masterpw string) error {
	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	w := zip.NewWriter(f)
	b, err := w.Encrypt("entries", string(masterpw))
	if err != nil {
		return err
	}
	err = binary.Write(b, binary.LittleEndian, a.entries)
	if err != nil {
		return err
	}
	return w.Close()

}

func OpenPwArchive(path string, password string) (*PwArchive, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	var pwa = new(PwArchive)
	pwa.entries = make([]AccountDetail, len(r.File))
	for _, f := range r.File {
		f.SetPassword(password)
		b, err := f.Open()
		if err != nil {
			return nil, err
		}

		err = binary.Read(b, binary.LittleEndian, pwa.entries)
		if err != nil {
			return nil, err
		}

	}

	return pwa, nil
}

// Search searches the fields of entries for strings that match any of the terms
// present in match (which should be a space seperated list), and returns a
// possibly zero length slice containing every AccountDetail which contained
// one of the target strings.
func (a *PwArchive) GlobSearch(match string) []AccountDetail {
	if a.Decrypted {
		return nil
	}
	var ad []AccountDetail
	for _, t := range strings.Split(match, " ") {
		for _, e := range a.entries {
			switch t {
			case e.Name, e.Password, e.LoginURL, e.Provider, e.ProviderDomain:
				ad = append(ad, e)
			}
		}
	}
	return ad
}

type AccountDetail struct {
	// The login name for this account
	Name string

	// The generated account password.
	// NOTE: this should be the actual password string, and not a hash of
	// it.
	Password string

	//Name of the provider
	Provider string

	// The providers domain addr
	ProviderDomain string

	// Optional full url to the providers login form.
	LoginURL string
}

// newAccountDetail creates an AccountDetail object using the arguments. If pw
// is "", a random one will be generated using DefaultRandPasswordLen.
func newAccountDetail(login string, pw, provider, domain, url string) AccountDetail {

}
