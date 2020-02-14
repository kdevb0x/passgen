package main

import (
	"fmt"
	"math/rand"
	_ "strconv"
	"strings"
	"time"
)

const (
	verifyL = "abcdefghijklmnopqrstuvwxyz"
	verifyU = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	verifyN = "1234567890"
	verifyS = "!@#$%&*"
)

var constraints = make(map[string]bool)

// generatePool builds a pool of valid charaters based on the given
// constraints, then returns it, ensuring it contains only valid characters.
func generatePool(constraints map[string]bool) []rune {
	var masterInclude []rune // 26+26+10+8
	for k, v := range constraints {
		if v == true {
			switch k {
			case "symbol":
				// 33, 35-38, 42, and 64
				for i := 35; i < 39; i++ {
					masterInclude = append(masterInclude, rune(i))
				}
				masterInclude = append(masterInclude, 33, 42, 64)
				continue
			case "number":
				// 48-57
				// TODO: change temp vars for each case.
				for i := 48; i < 58; i++ {
					masterInclude = append(masterInclude, rune(i))
				}
				continue
			case "upper":
				// 65-90
				for i := 65; i < 91; i++ {
					masterInclude = append(masterInclude, rune(i))
				}
				continue
			case "lower":
				// include 97..122 (a-z)
				for i := 97; i < 123; i++ {
					masterInclude = append(masterInclude, rune(i))
				}
				continue
			}
		}
	}
	return masterInclude
}

// checkRegen verifies a []string contains at least 1 char from every character class
// this function returns true if it DOES NOT contain, and needs to be regenerated.
func checkRegen(passStr []rune, classes map[string]bool) (regen bool) {
	for cls, v := range classes { // check what classes should be included
		if v == true {

			// for _, val := range passStr {
			var p = string(passStr)
			switch cls {
			case "lower":
				if !strings.ContainsAny(p, verifyL) {
					regen = true
				}
			case "upper":
				if !strings.ContainsAny(p, verifyL) {
					regen = true
				}
			case "number":
				if !strings.ContainsAny(p, verifyL) {
					regen = true
				}
			case "symbol":
				if !strings.ContainsAny(p, verifyL) {
					regen = true
				}
			}
		}

	}
	regen = false
	return
}
func generateChars(include []rune, length int) []rune {
	var passStr = make([]rune, length)
	rand.Seed(time.Now().UnixNano())

again:
	if len(include) > 0 {
		for i := 0; i < length; i++ {
			j := rand.Intn(len(include) - 1)
			passStr = append(passStr, include[j])
		}
	}

	// the function checkRegen checks to make sure only requested classess
	// were generated (invariant of the `verify` function).
	if regen := checkRegen(passStr, constraints); regen {

		// if not, generate another string and check again.
		goto again
	}

	return passStr[:]
}

func buildString(from []rune, verifyIt bool) (string, error) {
	var builder strings.Builder
	var s string // to hold builder.String()

	for _, char := range from {
		_, err := builder.WriteString(string(char))
		if err != nil {
			return "", err
		}
	}
	s = builder.String()
	if verifyIt { /* so named because verify is shadowed */
		if !verify(s) {
			return s, fmt.Errorf("unable to verify the content\n")
		}
	}
	return s, nil
}

// verify the contents of s; it returns true if s WAS verified, and false if it
// WAS NOT verified.
func verify(s string) bool {
	for _, f := range includeFlags {
		switch f {
		case 'l':
			if !strings.ContainsAny(s, verifyL) {
				return false
			}
		case 'n':
			if !strings.ContainsAny(s, verifyN) {
				return false
			}
		case 's':
			if !strings.ContainsAny(s, verifyS) {
				return false
			}
		case 'u':
			if !strings.ContainsAny(s, verifyU) {
				return false
			}
		default:
			// ignore it
		}
	}
	return true
}
