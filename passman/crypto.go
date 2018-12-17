package passman

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

)

type indexer func(*base64.Encoding) index

type index struct {
	indexNumber int
	references struct{
		username string `yaml:`
	}
}
type PassDataBase struct {
	map[indexer]*record
}
