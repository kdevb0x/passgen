// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"github.com/awnumar/memguard"
)

const defaultPassLen = 16
const defaultConstraints = "luns"

type fileHandler struct {
	f *memguard.Stream
}

func NewFileHandler() *fileHandler {
	return &fileHandler{memguard.NewStream()}
}

func (h *fileHandler) Add(provider, email, password string) error {
	if password == "" {
		password = genPassword(defaultPassLen, defaultConstraints)
	}
}
