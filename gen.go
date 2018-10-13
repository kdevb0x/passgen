package main

import (
	crand "crypto/rand"
	"encoding"
	"fmt"
	"time"
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

// Random returns
func randomL() []byte {
	lower := alphaL
	for i := rand.Seed(time.Now().UnixNano()
	_, err := crand.Read([]byte(lower))
}
