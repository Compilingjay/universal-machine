package um

import (
	"io"
	"log"
	"os"
	"slices"
)

func (um *UM) ConditionalMove(a, b, c uint32) {
	if um.Regs[c] != 0 {
		um.Regs[a] = um.Regs[b]
	}
}

func (um *UM) ArrayIndex(a, b, c uint32) error {
	um.Regs[a] = um.Arrs[um.Regs[b]][um.Regs[c]]
	return nil // possible out of bounds access, accessing abandoned array - Regs[c] >= len(Arr[Reg[b]])
}

func (um *UM) ArrayAmend(a, b, c uint32) error {
	um.Arrs[um.Regs[a]][um.Regs[b]] = um.Regs[c]
	return nil // possible out of bounds access, accessing abandoned array - Regs[b] >= len(Arr[Reg[a]])
}

func (um *UM) Addition(a, b, c uint32) {
	um.Regs[a] = uint32((uint64(um.Regs[b]) + uint64(um.Regs[c])) % ArithMax)
}

func (um *UM) Multiplication(a, b, c uint32) {
	um.Regs[a] = uint32((uint64(um.Regs[b]) * uint64(um.Regs[c])) % ArithMax)
}

func (um *UM) Division(a, b, c uint32) error {
	um.Regs[a] = um.Regs[b] / um.Regs[c]
	return nil // possible division by 0 - Regs[c] == 0
}

func (um *UM) Notand(a, b, c uint32) {
	um.Regs[a] = ^(um.Regs[b] & um.Regs[c])
}

func (*UM) HaltProgram() {
	os.Exit(0)
}

func (um *UM) ArrayAllocate(b, c uint32) {
	i, err := um.FreedArrs.pop()
	if err != nil {
		um.Arrs = append(um.Arrs, make([]uint32, um.Regs[c]))
		um.Regs[b] = uint32(len(um.Arrs) - 1)
		return
	}

	um.Arrs[i] = make([]uint32, um.Regs[c])
	um.Regs[b] = i
}

func (um *UM) ArrayAbandon(c uint32) error {
	um.FreedArrs.push(um.Regs[c])
	um.Arrs[um.Regs[c]] = nil
	return nil // possible access out of bounds - Regs[c] >= len(Arr[Reg[c]])
}

func (um *UM) Output(c uint32) error {
	um.IO.Writer.WriteByte(uint8(um.Regs[c]))
	um.IO.Writer.Flush()
	return nil // possible print of invalid char, Reg[c] > 255
}

func (um *UM) Input(c uint32) error {
	_, err := um.IO.Reader.Read(um.InChar[:])
	if err != nil {
		if err == io.EOF {
			um.Regs[c] = 0xffffffff
			return nil
		}
		log.Fatal(err.Error())
	}

	um.Regs[c] = uint32(um.InChar[0])
	return nil
}

func (um *UM) LoadProgram(b, c uint32) error {
	if um.Regs[b] != 0 {
		um.Arrs[0] = slices.Clone(um.Arrs[um.Regs[b]])
		um.ProgEnd = uint32(len(um.Arrs[0]))
	}

	um.PC = um.Regs[c]
	um.PC--

	return nil // possible abandoned array access, out of bounds array access - Reg[c] > len(Arr[Regs[b]])
}

func (um *UM) Orthography(instr uint32) {
	a := (instr >> 25) & 0b111
	um.Regs[a] = instr & 0x01ffffff
}
