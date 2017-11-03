package tools

func byteDistance(b1 byte, b2 byte) int {
	diff := b1 ^ b2
	d := 0
	for diff != 0 {
		diff -= diff & (^diff + 1)
		d += 1
	}
	return d
}

// Number of differing BITS
func EditDistance(s1 string, s2 string) int {
	diff := 0
	for i := 0; i < len(s1) && i < len(s2); i++ {
		diff += byteDistance(byte(s1[i]), byte(s2[i]))
	}
	return diff
}
