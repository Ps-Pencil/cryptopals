package tools

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hexString string) (string, error) {
	b, err := hex.DecodeString(hexString)
	if err != nil {
		return "", fmt.Errorf("Malformated hex string: %v", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
