package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"

	"github.com/PaloAltoNetworks/pan-saas-cli/SaaS-CLI/pkg/secret"
)

// Decrypt: returns a cleartext string
func Decrypt() string {

	// A 32 byte array used as the key; used for both encryption and decryption
	key := secret.Secret()

	cipherTxt, err := ioutil.ReadFile("./key.data")
	if err != nil {
		log.Fatal("Unable to read fw.json", err)
	}

	// generate a new aes cipher using the 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Unable to create AES Cipher", err)
	}

	// gcm is a module of operation for symmetric key
	// cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal("Unable to create GCM Cipher", err)
	}

	//Get the nounce size
	nounceSize := gcm.NonceSize()
	if len(cipherTxt) < nounceSize {
		log.Fatal("CipherText is smaller than the nounce")
	}

	//Extract nounce from the encrypted data
	nounce, cipherTxt := cipherTxt[:nounceSize], cipherTxt[nounceSize:]

	//Decrypt the data
	plainTxt, err := gcm.Open(nil, nounce, cipherTxt, nil)
	if err != nil {
		log.Fatal("GCM unable to open CiperText", err)
	}

	return string(plainTxt)
}
