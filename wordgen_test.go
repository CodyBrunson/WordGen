package main

import (
	"strings"
	"testing"
)

func TestWordGen_Generate(t *testing.T) {
	want := 3
	wordGen := NewWordGen()
	wordGen.Separator = " "
	words := strings.Fields(wordGen.Generate())
	if want != len(words) {
		t.Fatalf("`words = %q, expected length %v, actual length %v", words, want, len(words))
	}
}
