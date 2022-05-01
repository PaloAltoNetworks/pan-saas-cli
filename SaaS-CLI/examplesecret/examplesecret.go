package examplesecret

// Secret: returns a 32 byte array used as the encryption and decryption key. This can be any 32byte array.
func Secret() []byte {
	return []byte("thisshouldbeany32bytearrayyoudig")
}