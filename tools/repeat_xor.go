package tools

import "encoding/hex"

// RepeatXor takes two unicode string and returns the hex-encoded repeat-xor of the data.
func RepeatXor(data, key string) string {
	resultBytes := make([]byte, 0, len(data))

	for i, b := range data {
		resultBytes = append(resultBytes, byte(b)^byte(key[i%len(key)]))
	}
	return hex.EncodeToString(resultBytes)
}
