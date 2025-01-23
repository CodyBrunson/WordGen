package main

import (
	_ "embed"
	"io"
	"math/rand"
	"os"
	"strings"
)

type WordGen struct {
	Count     int
	Separator string
	Words     []string
}

func NewWordGen() WordGen {
	w := WordGen{}
	w.Count = 2
	w.Separator = "-"
	w.Words = readWordDict()
	return w
}

func (wg *WordGen) Generate() string {
	var pieces []string
	for i := 0; i <= wg.Count; i++ {
		pieces = append(pieces, wg.Words[rand.Int()%len(wg.Words)])
	}
	return strings.Join(pieces, wg.Separator)
}

func readWordDict() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(bytes), "\n")
	return words
}
