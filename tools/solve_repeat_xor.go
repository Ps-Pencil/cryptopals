package tools

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

var (
	Candidates = 2
)

type Sample struct {
	KeySize  int
	Distance float64
}

type Samples []Sample

func (s Samples) Len() int {
	return len(s)
}

func (s Samples) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Samples) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}

// Returns the most probable key sizes.
func findKeySizes(ct []byte) []int {
	result := make([]Sample, 0, 40)
	for keySize := 2; keySize <= 40; keySize++ {
		sampleBlocks := len(ct) / keySize
		if len(ct) < keySize*sampleBlocks {
			break
		}
		totalDist := 0.0
		for i := 0; i < sampleBlocks-1; i++ {
			d := EditDistance(string(ct[i*keySize:(i+1)*keySize]), string(ct[(i+1)*keySize:(i+2)*keySize]))
			totalDist += float64(d)
		}
		// Find average edit distance.
		totalDist /= float64(sampleBlocks - 1)

		// Normalise
		totalDist /= float64(keySize)

		result = append(result, Sample{
			KeySize:  keySize,
			Distance: totalDist,
		})
	}

	sort.Sort(Samples(result))

	keySizes := make([]int, 0, Candidates)
	for _, s := range result[:Candidates] {
		keySizes = append(keySizes, s.KeySize)
	}
	return keySizes
}

func transposeBytes(m [][]byte) [][]byte {
	result := [][]byte{}
	for i := 0; i < len(m[0]); i++ {
		result = append(result, []byte{})
		for _, b := range m {
			if i >= len(b) {
				continue
			}
			result[len(result)-1] = append(result[len(result)-1], b[i])
		}
	}
	return result
}

func transposeStrings(m []string) []string {
	result := []string{}
	for i := 0; i < len(m[0]); i++ {
		result = append(result, "")
		for _, b := range m {
			if i >= len(b) {
				continue
			}
			result[len(result)-1] += string(b[i])
		}
	}
	return result
}

func SolveRepeatXor(raw []byte) (string, error) {
	keySizes := findKeySizes(raw)

	bestScore := 0.0
	ans := ""

	for _, size := range keySizes {
		// Split into keysize blocks
		blocks := make([][]byte, 0, 0)
		for i := 0; i*size < len(raw); i += 1 {
			end := (i + 1) * size
			if end > len(raw) {
				end = len(raw)
			}
			blocks = append(blocks, raw[i*size:end])
		}

		// Transpose. So all first characters of all blocks become one block .. etc.
		chunks := transposeBytes(blocks)

		totScore := 0.0
		pts := []string{}
		for _, c := range chunks {
			pt, score, err := SingleByteXorDecipher(hex.EncodeToString(c))
			if err != nil {
				return "", fmt.Errorf("error SingleByteXorDecipher: %v.", err)
			}
			totScore += score
			pts = append(pts, pt)
		}
		totScore /= float64(len(chunks))
		if totScore > bestScore {
			bestScore = totScore
			ans = strings.Join(transposeStrings(pts), "")
		}
	}
	return ans, nil
}
