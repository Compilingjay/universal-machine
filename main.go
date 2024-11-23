package main

import (
	"log"
	"os"

	. "um"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./program <filename>")
	}

	instructions, err := getInstructionsFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	um := NewUM()
	LoadInstructions(&um, instructions)

	Run(&um)
}
