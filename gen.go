package main

import (
	"builtin"
	crand "crypto/rand"
	"encoding"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

const (
	alphaU = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaL = "abcdefghijklmnopqrstuvwxyz"
	nums   = "1234567890"
	symbs  = ".,/?!@#$%^&*()-+="
)

var (
	rgen = rand.New(crand)
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

// randomL returns a single lowercase english letter in the range of a-z
func randomL() []byte {
	lower := strings.NewReader(alphaL)

	var runelen, _, _ = lower.ReadRune()
	var runesize = unsafe.Sizeof(runelen)
	newchar := make([]byte, builtin.IntegerType(runesize))

	var bgint big.Int
	bgint.SetInt64(len(alphaL))

	gen := crand.Int(&bgint)
	newchar = append(newchar, []byte(alphaL[gen.Int64()]))

	return newchar

}

// randomU returns a single uppercase english letter in the range of A-Z
func randomU() []byte {
	upper := strings.NewReader(alphaU)

	var runelen, _, _ = lower.ReadRune()
	var runesize = unsafe.Sizeof(runelen)
	newchar := make([]byte, builtin.IntegerType(runesize))

	var bgint big.Int
	bgint.SetInt64(len(alphaU))

	gen := crand.Int(&bgint)
	newchar = append(newchar, []byte(alphaU[gen.Int64()]))

	return newchar

}
