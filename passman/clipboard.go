package passman

import (
	cb "github.com/atotto/clipboard"
)

// Pass is an abstract type for passing around password strings
type Pass struct {
	data []byte
}

func (s *Pass) String() string {
	return string(s.data)
}

// WriteToClipboard writes a user's password string to clipboard for pasting into password a field
func (s *Pass) WriteToClipboard() error {
	if err := cb.WriteAll(s.String()); err != nil {
		return err
	}
	return nil

}
