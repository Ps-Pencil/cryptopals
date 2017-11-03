package tools

import (
	"encoding/hex"
	"fmt"
)

func Xor(s1 string, s2 string) (string, error) {
	b1, err := hex.DecodeString(s1)

	if len(s1) != len(s2) {
		return "", fmt.Errorf("Different length strings %q and %q!", s1, s2)
	}

	if err != nil {
		return "", fmt.Errorf("Error decoding hex string %q: %v", s1, err)
	}

	b2, err := hex.DecodeString(s2)
	if err != nil {
		return "", fmt.Errorf("Error decoding hex string %q: %v", s2, err)
	}

	ans := make([]byte, 0, len(b1))

	for i := 0; i < len(b1); i++ {
		ans = append(ans, b1[i]^b2[i])
	}

	return hex.EncodeToString(ans), nil
}
