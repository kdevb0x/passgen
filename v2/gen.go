package main

import (
	"math/rand"
	"strings"
	"unicode/utf8"
)

var constraints map[string]bool

func checkConstraints(constraints map[string]bool) []int {
	var masterInclude = make([]int, 26+10+8)
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

			case "lower":
				// include 97..122 (a-z)

			}
		}
	}
}
func GenerateChars(include []int) []byte {

}
