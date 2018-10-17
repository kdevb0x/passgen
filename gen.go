package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var constraints = make(map[string]bool)

func checkConstraints(constraints map[string]bool) []int {
	var masterInclude = make([]int, 26+26+10+8)
	for k, v := range constraints {
		if v == true {
			switch k {
			case "symbol":
				// 33-38, 42, and 64
				for i := 33; i < 39; i++ {
					masterInclude = append(masterInclude, i)
				}
				masterInclude = append(masterInclude, 42, 64)
			case "number":
				// 48-57
				for i := 48; i < 58; i++ {
					masterInclude = append(masterInclude, i)
				}
			case "upper":
				// 65-90
				for i := 65; i < 91; i++ {
					masterInclude = append(masterInclude, i)
				}
			case "lower":
				// include 97..122 (a-z)
				for i := 97; i < 123; i++ {
					masterInclude = append(masterInclude, i)
				}
			}
		}
	}
	return masterInclude
}
func generateChars(include []int) []string {
	var passStr []string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		j := rand.Intn(len(include) - 1)
		passStr[i] = strconv.Itoa(include[j])
	}
	return passStr
}

func buildString(from []string) (string, error) {
	var builder strings.Builder
	for _, char := range from {
		_, err := builder.WriteString(char)
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
