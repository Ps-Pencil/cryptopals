package tools

import "testing"

func TestXor(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := Xor(input1, input2)

	if err != nil {
		t.Errorf("Error xor: %v.", err)
	}

	if got != want {
		t.Errorf("Xor got %q; want %q.", got, want)
	}
}
