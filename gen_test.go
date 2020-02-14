package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGeneratePool(t *testing.T) {
	for _, c := range cases {
		var inc = make(map[string]bool)

		for _, l := range c.includes {
			r, _ := utf8.DecodeRuneInString(string(l))
			switch r {
			case 'l':
				inc["lower"] = true
			case 'u':
				inc["upper"] = true
			case 'n':
				inc["number"] = true
			case 's':
				inc["symbol"] = true
			default:
				t.Logf("unrecognized character class %s\n", string(r))
				t.Fail()
			}
		}
		pool := generatePool(inc)
		for _, m := range c.includes {
			switch m {
			case 'l':
				if !strings.ContainsAny(string(pool), verifyL) {
					t.Fail()
				}
			case 'u':
				if !strings.ContainsAny(string(pool), verifyU) {
					t.Fail()
				}
			case 'n':
				if !strings.ContainsAny(string(pool), verifyN) {
					t.Fail()
				}
			case 's':
				if !strings.ContainsAny(string(pool), verifyS) {
					t.Fail()
				}
			}
		}

	}
}

func TestGenerateChars(t *testing.T) {
	for _, c := range cases {
		cp := genConstraints(t, c)
		pool := generatePool(cp)
		chars := generateChars(pool, c.length)
		if len(chars) != c.length {
			t.Fail()
		}
	}
}

func genConstraints(t *testing.T, from testcase) map[string]bool {
	var inc = make(map[string]bool)

	for _, l := range from.includes {
		r, _ := utf8.DecodeRuneInString(string(l))
		switch r {
		case 'l':
			inc["lower"] = true
		case 'u':
			inc["upper"] = true
		case 'n':
			inc["number"] = true
		case 's':
			inc["symbol"] = true
		default:
			t.Logf("unrecognized character class %s\n", string(r))
			t.Fail()
		}
	}
	return inc
}
