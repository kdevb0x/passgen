// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

const (
	yes bool = true
	no  bool = false
)

var cases = []testcase{
	{
		includes: "luns",
		length:   16,
		verify:   yes,
	}, {
		includes: "s",
		length:   12,
		verify:   no,
	}, {
		includes: "lun",
		length:   54,
		verify:   yes,
	}, {
		includes: "uns",
		length:   -1,
		verify:   no,
	},
}

type testcase struct {
	includes string
	length   int
	verify   bool
}

func TestVerifyIntegrationTest(t *testing.T) {
	// lets do a smoketest and two unit-tests; one with verify flag, one
	// without.
	err := verifySmokeTest(t)
	if err != nil {
		t.Errorf("failed smoke-test; err: %w\n", err)
		t.Fail()
	}

	//
}

func VerifyFlagUnitTest(t *testing.T) {
	for _, c := range cases {
		cons := genConstraints(t, c)
		p := generatePool(cons)
		g := generateChars(p, c.length)

		s, err := buildString(g, c.verify)
		if err != nil {
			t.Errorf("error building string: %w\n", err)
		}
		if c.verify {
			if !verify(s) {
				t.Fail()
			}
		}

	}
}

func noVerifySmokeTest(t *testing.T) error {
	cmd := exec.Command("passgen", "-l 16 -i s")
	s, err := cmd.Output()
	if err != nil {
		return err
	}
	// lets take a chance and generalize to say that if the length is as
	// requested, all is well.
	if len(s) != 16 {
		switch {
		case len(s) < 16:
			return fmt.Errorf("string is too short; %d is shorter than the requested length.", len(s))
		case len(s) > 16:
			return fmt.Errorf("string is too long; %d is longer than it should be.", len(s))
		}
	}
	return nil
}

func verifySmokeTest(t *testing.T) error {
	cmd := exec.Command("passgen", "-l 16 -i s -v")
	s, err := cmd.Output()
	if err != nil {
		return err
	}
	if !bytes.ContainsAny(s, "!@#$%^&*") {
		t.Fail()

		// BUG: this line throws error: Errorf format %^ has unknown verb ^
		// return fmt.Errorf(`Error, string does not contain at least one [!@#$%^&*] symbol\n`)
		return errors.New("error: string does not contain at least one [!@#$%^&*] symbol\n")
	}
	return nil
}
