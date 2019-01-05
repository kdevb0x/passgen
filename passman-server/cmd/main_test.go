// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"

	env "github.com/joho/godotenv"
)

var _ = env.Load(nil)

func TestGetHostname(t *testing.T) {
	hostname, err := getHostname() // err for .env file not found, etc.
	if err != nil {
		fmt.Errorf("error while fetching $HOSTNAME: %s, received $HOSTNAME value: %s\n", err, hostname)
	}

}
