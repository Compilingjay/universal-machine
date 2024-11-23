package main

import (
	"bytes"
	"encoding/binary"
	"os"
)

func getInstructionsFromFile(filepath string) ([]uint32, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return []uint32{}, err
	}

	instructions := make([]uint32, len(b)/4)
	err = binary.Read(bytes.NewReader(b), binary.BigEndian, instructions)
	if err != nil {
		return []uint32{}, err
	}

	return instructions, nil
}
