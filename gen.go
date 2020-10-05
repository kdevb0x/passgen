package main

import (
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

type Constaints map[string]bool

func (cn *Constaints) Check() []int {
	if len(cn) > 0 {
		return checkConstraints(cn)
	}
	return nil
} 

func (cn *Constaints) GenerateChars(include []int) []string {
	return generateChars(include)
}


func (cn *Constaints) String() string {
	if cn != nil || len(cn) > 0 {
		s, err := buildString(cn.GenerateChars(cn.Check()))
		if err != nil {
			panic(err)
		}
		return s
	}
}

var constraints = make(map[string]bool)

func checkConstraints(constraints map[string]bool) []int {
	var masterInclude []int // 26+26+10+8
	for k, v := range constraints {
		if v == true {
			switch k {
			case "symbol":
				// 33, 35-38, 42, and 64
				for i := 35; i < 39; i++ {
					masterInclude = append(masterInclude, i)
				}
				masterInclude = append(masterInclude, 33, 42, 64)
				break
			case "number":
				// 48-57
				// TODO: change temp vars for each case.
				for i := 48; i < 58; i++ {
					masterInclude = append(masterInclude, i)
				}
				break
			case "upper":
				// 65-90
				for i := 65; i < 91; i++ {
					masterInclude = append(masterInclude, i)
				}
				break
			case "lower":
				// include 97..122 (a-z)
				for i := 97; i < 123; i++ {
					masterInclude = append(masterInclude, i)
				}
				break
			}
		}
	}
	return masterInclude
}

// checkRegen verifies a []string contains at least 1 char from every character class
func checkRegen(passStr []string, classes *map[string]bool) (regen bool) {
	for cls, v := range *classes { // check what classes should be included
		if v == true {

			for _, val := range passStr {
				switch cls {
				case "lower":
					if ok := strings.ContainsAny(val, verifyL); !ok {
						regen = true
					}
				case "upper":
					if ok := strings.ContainsAny(val, verifyL); !ok {
						regen = true
					}
				case "number":
					if ok := strings.ContainsAny(val, verifyL); !ok {
						regen = true
					}
				case "symbol":
					if ok := strings.ContainsAny(val, verifyL); !ok {
						regen = true
					}
				}
			}
		}

	}
	regen = false
	return
}
func generateChars(include []int) []string {
	var passStr []string
	rand.Seed(time.Now().UnixNano())

again:
	if len(include) > 0 {
		for i := 0; i < length; i++ {
			j := rand.Intn(len(include) - 1)
			passStr[i] = string(include[j])
		}
	}

	// verify passStr includes at least 1 of all character classes
	if regen := checkRegen(passStr, &constraints); regen {
		goto again
	}

	return passStr
}

func buildString(from []string) (string, error) {
	var builder strings.Builder
	for _, char := range from {
		_, err := builder.WriteString(string(char))
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
