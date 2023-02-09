package main

import (
	"os"
)

//go:generate go run . ../../multistring_gen.go

func main() {
	if len(os.Args) < 2 {
		panic("must provide destination path")
	}

	bc := generationContext{}
	bc.generate(os.Args[1])
	if bc.err != nil {
		panic(bc.err.Error())
	}
}
