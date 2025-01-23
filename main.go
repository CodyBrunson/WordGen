package main

import (
	_ "embed"
	"fmt"
)

func main() {
	wordGen := NewWordGen()
	fmt.Println(wordGen.Generate())
}
