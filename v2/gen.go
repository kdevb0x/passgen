package main

import (
	"math/rand"
	"strings"
	"time"
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
func generateChars(include []int) ([]byte, error) {
	var passStr []byte
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= length; i++ {
		j := rand.Intn(findHighest(include))
	}
}

func findHighest(nums []int) int {
	var high int
	for _, m := range nums {
		if m > high {
			high = m
		}
	}
	return high

}

func findLowest(nums []int) int {
	var low int
	for _, m := range nums {
		if m < low {
			low = m
		}
	}
	return low
}
