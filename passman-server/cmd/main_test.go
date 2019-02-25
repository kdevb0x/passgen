// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package main

import (
	"testing"

	env "github.com/joho/godotenv"
)

var _ = env.Load("")

func TestDBConnect(t *testing.T) {
	const LOCALDB = "couchbase://localhost"
	if err := dbConnect(LOCALDB); err != nil {
		t.Error(err)
		t.Fail()
	}
	
}

