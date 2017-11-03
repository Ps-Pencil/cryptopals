package tools

import (
	"crypto/aes"
	"fmt"
)

func AesEcbDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher: %v", err)
	}

	if len(ciphertext)%len(key) != 0 {
		return nil, fmt.Errorf("ciphertext length is not a multiple of key length")
	}
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i*len(key) < len(ciphertext); i++ {
		block.Decrypt(plaintext[i*len(key):], ciphertext[i*len(key):])
	}
	return plaintext, nil
}
