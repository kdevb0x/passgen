package passman

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"hash"
	"io"
	"log"
	"os"

	"github.com/gonuts/binary"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/pbkdf2"
)

func HashSHA512(key string) string {
	hsh := sha512.New()
	hsh.Write([]byte(key))
	return base58.Encode(hsh.Sum(nil))
}

func encryptAES(data []byte, pass PassString) []byte {
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

func decryptAES(data []byte, pass PassString) []byte {

}