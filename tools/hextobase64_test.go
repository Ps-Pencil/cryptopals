package tools

import "testing"

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	got, err := HexToBase64(input)
	if err != nil {
		t.Errorf("Error HexToBase64: %v.", err)
	}
	if got != want {
		t.Errorf("HexToBase64 got %q; want %q.", got, want)
	}
}
