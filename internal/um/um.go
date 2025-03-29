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
	CMov uint32 = iota
	ArrIndex
	ArrAmend
	Add
	Mult
	Div
	Nand
	Halt
	AllocArr
	AbandonArr
	Out
	In
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
	InChar    [1]uint8
	IO        *bufio.ReadWriter
}

func NewUM() UM {
	return UM{
		Regs:      [8]uint32{0, 0, 0, 0, 0, 0, 0, 0},
		Arrs:      make([][]uint32, 1),
		FreedArrs: Queue{},
		PC:        0,
		ProgEnd:   0,
		InChar:    [1]uint8{0},
		IO: bufio.NewReadWriter(
			bufio.NewReader(os.Stdin),
			bufio.NewWriter(os.Stdout),
		),
	}
}

func (um *UM) LoadInstructions(instructions []uint32) {
	um.Arrs[0] = slices.Clone(instructions)
	um.ProgEnd = uint32(len(um.Arrs[0]))
}

func (um *UM) Run() {
	for um.PC < um.ProgEnd {
		instr := um.Arrs[0][um.PC]
		op := (instr >> 28) & 0b1111
		a := (instr >> 6) & 0b111
		b := (instr >> 3) & 0b111
		c := (instr & 0b111)

		switch op {
		case CMov:
			um.ConditionalMove(a, b, c)
		case ArrIndex:
			_ = um.ArrayIndex(a, b, c)
		case ArrAmend:
			um.ArrayAmend(a, b, c)
		case Add:
			um.Addition(a, b, c)
		case Mult:
			um.Multiplication(a, b, c)
		case Div:
			_ = um.Division(a, b, c)
		case Nand:
			um.Notand(a, b, c)
		case Halt:
			um.HaltProgram()
		case AllocArr:
			um.ArrayAllocate(b, c)
		case AbandonArr:
			_ = um.ArrayAbandon(c)
		case Out:
			_ = um.Output(c)
		case In:
			_ = um.Input(c)
		case LoadProg:
			_ = um.LoadProgram(b, c)
		case Orth:
			um.Orthography(instr)
		default:
			log.Fatalf("invalid opcode: %d", op)
		}

		um.PC++
	}
}
