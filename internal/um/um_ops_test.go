package um_test

import (
	"testing"

	. "um"
)

func TestConditionalMove(t *testing.T) {
	um := NewUM()

	var a, b, c uint32 = 0, 1, 2
	var v1, v2, v3 uint32 = 6, 42, 0
	um.Regs[0] = v1
	um.Regs[1] = v2
	um.Regs[2] = v3

	// no conditional move
	ConditionalMove(&um, a, b, c)
	if um.Regs[a] != v1 {
		t.Errorf("denied conditional move error, expected: %d, got: %d", v1, um.Regs[a])
	}

	// conditional move
	um.Regs[c] = 1
	ConditionalMove(&um, a, b, c)
	if um.Regs[a] != v2 {
		t.Errorf("permitted conditional move failed, expected: %d, got: %d", v2, um.Regs[a])
	}
}

func TestArithmetic(t *testing.T) {
	um := NewUM()
	var a, b, c uint32 = 0, 1, 2
	var expected uint32 = 0x00000000
	um.Regs[b] = 0xffffffff
	um.Regs[c] = 0x00000001
	Addition(&um, a, b, c)
	if um.Regs[a] != expected {
		t.Errorf("addition failed, expected: %d, got: %d", expected, um.Regs[a])
	}

	expected = 0x00000000
	um.Regs[b] = 0x80000000
	um.Regs[c] = 0x00000002
	Multiplication(&um, a, b, c)
	if um.Regs[a] != expected {
		t.Errorf("multiplication failed, expected: %d, got: %d", expected, um.Regs[a])
	}
}

func TestArrayAllocUseAndAbandon(t *testing.T) {
	um := NewUM()

	var a, b, c uint32 = 3, 4, 5
	var arrSize uint32 = 16
	um.Regs[c] = arrSize

	ArrayAllocate(&um, b, c)
	if um.Regs[b] != 1 {
		t.Errorf("array identifier not returned, expected: %d, got: %d", 1, um.Regs[b])
	}
	for _, v := range um.FreedArrs {
		if v == um.Regs[b] {
			t.Errorf("allocated array is not set as used, arr: %d", um.Regs[b])
		}
	}

	actualArrSize := len(um.Arrs[um.Regs[b]])
	if actualArrSize != 16 {
		t.Errorf("incorrect array size, expected: %d, got: %d", arrSize, actualArrSize)
	}

	um.Regs[a] = um.Regs[b] // move array identifier to register a
	var offset, v1 uint32 = 8, 61
	um.Regs[b] = offset
	um.Regs[c] = v1

	_ = ArrayAmend(&um, a, b, c)
	actualValue := um.Arrs[um.Regs[a]][offset]
	if actualValue != v1 {
		t.Errorf("invalid array amendment, expected: %d, got: %d", v1, actualValue)
	}

	um.Regs[b] = um.Regs[a] // move array identifier to register b
	um.Regs[c] = offset
	_ = ArrayIndex(&um, a, b, c)
	if um.Regs[a] != v1 {
		t.Errorf("array index not returning right value, expected: %d, got: %d", v1, um.Regs[a])
	}

	um.Regs[c] = um.Regs[b] // move array identifier to register c
	_ = ArrayAbandon(&um, c)
	for _, v := range um.FreedArrs {
		if v == um.Regs[c] {
			return
		}
	}

	t.Errorf("array abandonment not executed, arr: %d", um.Regs[c])
}

func TestNand(t *testing.T) {
	um := NewUM()

	var a, b, c uint32 = 1, 6, 7
	var v1, v2 uint32 = 0xffff00ff, 0xffff0f30
	var expected uint32 = 0x0000ffcf
	um.Regs[b] = v1
	um.Regs[c] = v2

	Notand(&um, a, b, c)
	if um.Regs[a] != expected {
		t.Errorf("nand result not valid, expected: %d, got: %d", expected, um.Regs[a])
	}
}

func TestOrthography(t *testing.T) {
	um := NewUM()

	var instr uint32 = 0xdf0000ff // op - 13, reg - 7, v == expected
	var expected uint32 = 0x010000ff
	Orthography(&um, instr)
	if um.Regs[7] != expected {
		t.Errorf("orthography result not valid, expected: %d, got: %d", expected, um.Regs[7])
	}
}
