package main

import (
	crand "crypto/rand"
	"encoding"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

const (
	alphaU = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaL = "abcdefghijklmnopqrstuvwxyz"
	nums   = "1234567890"
	symbs  = ".,/?!@#$%^&*()-+="
)

type constraints string

var (
	constr = make(map[constraints]bool, 100)
	rgen   = rand.New(crand)
)

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: not equal lengths")
	}

	result := make([]byte, len(a))
	for i, j := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

// Random returns
func randomL() int {
	lower := alphaL
	randr := crand.Read([]byte(lower))

}
