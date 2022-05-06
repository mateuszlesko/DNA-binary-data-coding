package main

import "testing"

func TestCodeToDNA(t *testing.T) {
	dna := codeToDNA(66)
	if dna != "AGAAC" {
		t.Errorf("DNA code was incorrect, got: %s, want: %s.", dna, "AGAAC")
	}
}
