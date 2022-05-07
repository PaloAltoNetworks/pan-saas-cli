package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"

	"github.com/zackmacharia/PANOS-GOLANG/pkg/secret"
)

// Encrypt: takes a cleartext string and encrypts it into an array of bytes
func Encrypt(ctxt string) []byte {

	// A 32 byte array used as the key; used for both encryption and decryption
	key := secret.Secret() 

	text := []byte(ctxt)

	// generate a new aes cipher using the 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Unable to create AES Cipher", err)
	}

	// gcm is a module of operation for symmetric key
	// cryptographic block ciphers - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal("Unable to create GCM Cipher")
	}

	//Create a new byte array of the same size as the nounce
	// then pass it to Seal
	nounce := make([]byte, gcm.NonceSize())
	
	//Populate nounce with a cryptographically random sequence
	if _, err = io.ReadFull(rand.Reader, nounce); err != nil {
		log.Fatal("Unable to read random nounce")
	}

	//Seal encrypts the data
	return gcm.Seal(nounce, nounce, text, nil)
}
