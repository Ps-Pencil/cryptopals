package tools

import "testing"

func TestRepeatXor(t *testing.T) {
	table := []struct {
		Input string
		Key   string
		Want  string
	}{
		{
			Input: `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`,
			Key:  "ICE",
			Want: "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272" + "a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f",
		},
	}

	for _, c := range table {
		got := RepeatXor(c.Input, c.Key)
		if got != c.Want {
			t.Errorf("RepeatXor: \ngot  %q; \nwant %q.", got, c.Want)
		}
	}
}
