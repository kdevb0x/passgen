// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

// Refs holds the references to user accounts that a PW is associated with
type Refs map[string][]string

type PW struct {
	PWHash     []byte // PASSWORD MUST BE STORED HASHED
	References *Refs
}

// NewPW returns a pointer to a PW which holds a hashed password
func NewPW(hashedPass []byte) *PW {
	pw := &PW{PWHash: hashedPass}
	return pw
}
