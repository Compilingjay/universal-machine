package um

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
)

// type Opcode uint32

const (
	CondMove uint32 = iota
	ArrIndex
	ArrAmend
	Add
	Mult
	Divide
	Nand
	Halt
	AllocArr
	AbandonArr
	Outp
	Inp
	LoadProg
	Orth
)

const (
	ArithMax = math.MaxUint32 + 1
)

type UM struct {
	Regs      [8]uint32
	Arrs      [][]uint32
	FreedArrs Queue
	PC        uint32
	ProgEnd   uint32
	InChar    []uint8
	IO        *bufio.ReadWriter
}

func NewUM() UM {
	return UM{
		Regs:      [8]uint32{0, 0, 0, 0, 0, 0, 0, 0},
		Arrs:      make([][]uint32, 1),
		FreedArrs: Queue{},
		PC:        0,
		ProgEnd:   0,
		InChar:    make([]uint8, 1),
		IO: bufio.NewReadWriter(
			bufio.NewReader(os.Stdin),
			bufio.NewWriter(os.Stdout),
		),
	}
}

func LoadInstructions(um *UM, instructions []uint32) {
	um.Arrs[0] = slices.Clone(instructions)
	um.ProgEnd = uint32(len(um.Arrs[0]))
}

func Run(um *UM) {
	for um.PC < um.ProgEnd {
		instr := um.Arrs[0][um.PC]
		op := (instr >> 28) & 0b1111
		a := (instr >> 6) & 0b111
		b := (instr >> 3) & 0b111
		c := (instr & 0b111)

		switch op {
		case CondMove:
			ConditionalMove(um, a, b, c)
		case ArrIndex:
			ArrayIndex(um, a, b, c)
		case ArrAmend:
			ArrayAmend(um, a, b, c)
		case Add:
			Addition(um, a, b, c)
		case Mult:
			Multiplication(um, a, b, c)
		case Divide:
			Division(um, a, b, c)
		case Nand:
			Notand(um, a, b, c)
		case Halt:
			HaltProgram()
		case AllocArr:
			ArrayAllocate(um, b, c)
		case AbandonArr:
			ArrayAbandon(um, c)
		case Outp:
			Output(um, c)
		case Inp:
			Input(um, c)
		case LoadProg:
			LoadProgram(um, b, c)
		case Orth:
			Orthography(um, instr)
		default:
			log.Fatalf("invalid opcode: %d", op)
		}

		um.PC++
	}
}
