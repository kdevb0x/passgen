package main

import (
	"bufio"
	"bytes"
	crand "crypto/rand"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"strings"
	"time"
	"unsafe"

	"github.com/codegangsta/cli"
)

const (
	classes = []string{0: "lowercase", 1: "uppercase", 2: "numbers", 3: "symbols"}
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "1234567890"
	symbols   = ".,/?!@#$%^&*()-+="
)

var (
	rgen = newPassGen()
)

type Generator interface {
	Generate()
	String() string
}

type passgen struct {
	length      int
	constraints map[string]bool
	buff        []byte
	reader      io.Reader
}

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
	lower := strings.NewReader(lowercase)

	var runelen, _, _ = lower.ReadRune()
	var runesize = unsafe.Sizeof(runelen)
	newchar := make([]byte, int(runesize))

	var bgint big.Int
	bgint.SetInt64(len(lowercase))

	gen := crand.Int(&bgint)
	newchar = append(newchar, []byte(lowercase[gen.Int64()]))

	return newchar

}

// randomU returns a single uppercase english letter in the range of A-Z
func randomU() []byte {
	upper := strings.NewReader(uppercase)

	var runelen, _, _ = lower.ReadRune()
	var runesize = unsafe.Sizeof(runelen)
	newchar := make([]byte, builtin.IntegerType(runesize))

	var bgint big.Int
	bgint.SetInt64(len(uppercase))

	gen := crand.Int(&bgint)
	newchar = append(newchar, []byte(uppercase[gen.Int64()]))

	return newchar

}

// NewGenerator creates an empty Generator interface
func NewGenerator() Generator {
	var gen = new(Generator)
	return gen
}

// newPassGen returns a pointer to a newly created passgen
func NewPassGen(strlen int) *passgen {
	var bf = make([]byte, 1024)
	gen := &passgen{
		length:      strlen,
		constraints: make(map[string]bool{classes[0]: true, classes[1]: false, classes[2]: false, classes[3]: false}),
		buff:        bf,
		reader:      bytes.NewBuffer(buff),
	}

	return gen
}

func (n *passgen) Generate() {
	var include []string
	for k, v := range n.constraints {
		if v == true {
			append(include, k)
		}
	}
	chanque := make([]<-chan []byte, len(include))
	for _, i := range include {
		switch i {
		case "lowercase":
			append(chanque, asyncRandomL())
			break
		case "uppercase":
			append(chanque, asyncRandomU())
			break
		case "numbers":
			append(chanque, asyncRandomN())
			break
		case "symbols":
			append(chanque, asyncRandomS())
			break
		default:
			fmt.Printf("warning no constraints: can't generate a password string from nothing!")
			return
		}
	}
}
