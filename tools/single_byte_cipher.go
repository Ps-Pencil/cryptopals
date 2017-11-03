package tools

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func ScoreEnglish(src []byte) float64 {
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ,.?-\"'"

	score := 0.0
	freq := map[byte]int{}
	maxFreq := 0
	var maxChar byte
	for _, b := range src {
		if strings.ContainsRune(validChars, rune(b)) {
			score += 1
		} else {
			score -= 1
		}
		freq[b] += 1
		if freq[b] > maxFreq {
			maxFreq = freq[b]
			maxChar = b
		}
	}

	// Some heuristic.
	if maxChar == byte('e') || maxChar == byte('E') {
		score *= 1.1
	}

	return score
}

// src is hex encoded.
func SingleByteXorDecipher(src string) (string, float64, error) {
	var ans string
	best := 0.0
	for k := 0; k < 256; k++ {
		xorPad := strings.Repeat(hex.EncodeToString([]byte{byte(k)}), len(src)/2)
		candidate, err := Xor(src, xorPad)
		if err != nil {
			return "", 0, fmt.Errorf("error xoring strings: %v", err)
		}
		decoded, err := hex.DecodeString(candidate)
		if err != nil {
			return "", 0, fmt.Errorf("error decoding string: %v", err)
		}
		score := ScoreEnglish(decoded)
		if score > best {
			best = score
			ans = string(decoded)
		}
	}
	return ans, best, nil
}
