package bignumbers

import (
	"math"
	"strings"
)

type BigNumber struct {
	blocks []Uint
}

func (bn *BigNumber) SetHex(hex string) error {
	stringBlocks := breakStringIntoBlocks(hex, 16)
	uintBlocks := make([]Uint, 0)
	for _, b := range stringBlocks {
		var u Uint
		if err := u.SetHex(b); err != nil {
			return err
		}
		uintBlocks = append(uintBlocks, u)
	}
	bn.blocks = uintBlocks

	return nil
}

func (bn *BigNumber) GetHex() (hex string) {
	for i, block := range bn.blocks {
		blockHex := block.GetHex()
		if i != len(bn.blocks)-1 {
			missingZerosCount := 16 - len(blockHex)
			var sb strings.Builder
			for i := 0; i < missingZerosCount; i++ {
				sb.WriteString("0")
			}
			sb.WriteString(blockHex)
			blockHex = sb.String()
		}
		hex = blockHex + hex
	}
	return hex
}

func (bn *BigNumber) Invert() (result BigNumber) {
	invertedBlocks := make([]Uint, len(bn.blocks))
	for i, block := range bn.blocks {
		invertedBlocks[i] = block.Invert()
	}
	result.blocks = invertedBlocks
	resultHex := result.GetHex()
	resultHex = resultHex[len(resultHex)-len(bn.GetHex()):]
	result.SetHex(resultHex)
	return
}

func binaryOperation(a, b BigNumber, operation func(Uint, Uint) uint64) (result BigNumber) {
	blocks := make([]Uint, int(math.Max(float64(len(a.blocks)), float64(len(b.blocks)))))
	for i, j := 0, 0; i < len(a.blocks) || j < len(b.blocks); i, j = i+1, j+1 {
		if i >= len(a.blocks) {
			blocks[i] = b.blocks[j]
		} else if j >= len(b.blocks) {
			blocks[i] = a.blocks[i]
		} else {
			var u Uint
			u.SetDecimal(operation(a.blocks[i], b.blocks[j]))
			blocks[i] = u
		}
	}
	result.blocks = blocks
	return
}

func (bn *BigNumber) XOR(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) uint64 { return u1.value ^ u2.value })
}

func (bn *BigNumber) AND(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) uint64 { return u1.value & u2.value })
}
