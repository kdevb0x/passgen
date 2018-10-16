package main

import (
	"math/rand"
)

var (
	classes = []string{0: "lowercase", 1: "uppercase", 2: "numbers", 3: "symbols"}
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "1234567890"
	symbols   = "!@#$%&*"
)

type Generator interface {
	Generate()
	String() string
}

type passgen struct {
	length      int
	constraints map[string]bool
	buff        []byte
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
func (n passgen) randomL() func() []byte {
	f := func() []byte {

		newchar := make([]byte, 1)
		rnum := rand.Intn(len(lowercase))
		newchar = append(newchar, lowercase[rnum-1])

		return newchar
	}
	return f
}

// randomU returns a single uppercase english letter in the range of A-Z
func (n passgen) randomU() func() []byte {
	f := func() []byte {

		newchar := make([]byte, 1)
		rnum := rand.Intn(len(uppercase))
		newchar = append(newchar, uppercase[rnum-1])

		return newchar
	}
	return f
}

func (n passgen) randomN() func() []byte {
	f := func() []byte {

		newchar := make([]byte, 1)
		rnum := rand.Intn(len(numbers))
		newchar = append(newchar, numbers[rnum-1])

		return newchar
	}
	return f
}

func (n passgen) randomS() func() []byte {
	f := func() []byte {

		newchar := make([]byte, 1)
		rnum := rand.Intn(len(symbols))
		newchar = append(newchar, symbols[rnum-1])

		return newchar
	}
	return f
}

// NewGenerator creates an empty Generator interface
func NewGenerator() Generator {
	var gen = new(Generator)
	return *gen
}

// newPassGen returns a pointer to a newly created passgen
func newPassGen(strlen int) *passgen {
	var bf = make([]byte, 1024)
	gen := &passgen{
		length:      strlen,
		constraints: constraint,
		buff:        bf,
	}

	return gen
}

func (n *passgen) Generate() {

	var include []string
	for k, v := range n.constraints {
		if v == true {
			include = append(include, k)
		}
	}
	var fnque []func() []byte

	for _, i := range include {
		switch i {
		case "lowercase":
			_ = append(fnque, n.randomL())
			break
		case "uppercase":
			_ = append(fnque, n.randomU())
			break
		case "numbers":
			_ = append(fnque, n.randomN())
			break
		case "symbols":
			_ = append(fnque, n.randomS())
			break
		default:

		}
	}
	// pick which type of character to generate at random
	for i, j := 0, rand.Intn(len(fnque)); i <= n.length; i++ {
		n.buff[i] = fnque[j]()[0]
	}
}

func (n passgen) String() string {
	var ln = len(n.buff)
	stb := string(n.buff[:ln])
	return stb

}
