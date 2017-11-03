package tools

import "testing"

func TestEditDistance(t *testing.T) {
	s1 := "this is a test"
	s2 := "wokka wokka!!!"

	got := EditDistance(s1, s2)
	want := 37
	if got != want {
		t.Errorf("EditDistance got %d; want %d.", got, want)
	}
}
