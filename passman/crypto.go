// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"io"
	"log"

	"github.com/mr-tron/base58"
)



func HashSHA512(key string) string {
	hsh := sha512.New()
	hsh.Write([]byte(key))
	return base58.Encode(hsh.Sum(nil))
}

func encryptAES(data []byte, pass Pass) []byte {
	blk, err := aes.NewCipher([]byte(HashSHA512(pass.String())))
	if err != nil {
		log.Fatalf("error creating encryption cipher: %s", err)
	}

	gal, err := cipher.NewGCM(blk)
	nonce := make([]byte, gal.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ctext := gal.Seal(nonce, nonce, data, nil)
	return ctext
}

func decryptAES(data []byte, pass Pass) []byte {

}
