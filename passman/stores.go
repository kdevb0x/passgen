package passman

import (
	"encoding"
	"fmt"
	toml "github.com/pelletier/go-toml"
	"io"
)

type Algorithm uint8

const (
	_ Algotithm = iota
	AES
	Pdkfd2
	Twofish
)

type SecretMarshaler interface {
	MarshalSecret(algo Algorithm, to io.Writer) error
	UnmarshalSecret(algo Algorithm) error
	fmt.Stringer
}

type secret struct {
	length    int
	data      *byte
	encrypted bool
}

func (s *secret) MarshalSecret(algo Algorithm, to io.Writer) error {
	switch algo {
	case AES:
		if s.encrypted == false {
			enc := encryptAES([]byte(*s.data[:s.length]))
		}
	}
}
