package bignumbers

import (
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
