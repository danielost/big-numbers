package bignumbers

import (
	"strings"
)

type BigNumber struct {
	blocks []Uint
}

func (bn *BigNumber) GetBlocks() []Uint {
	return bn.blocks
}

func (bn *BigNumber) SetHex(hex string) error {
	return bn.setValue(16, hex, func(s string) (Uint, error) {
		var u Uint
		if err := u.SetHex(s); err != nil {
			return Uint{0}, err
		}
		return u, nil
	})
}

func (bn *BigNumber) SetBinary(binary string) error {
	return bn.setValue(64, binary, func(s string) (Uint, error) {
		var u Uint
		if err := u.SetBinary(s); err != nil {
			return Uint{0}, err
		}
		return u, nil
	})
}

func (bn *BigNumber) setValue(blockSize int, value string, setter func(string) (Uint, error)) error {
	inputBlocks := breakStringIntoBlocks(value, blockSize)
	resultBlocks := make([]Uint, 0)
	for _, block := range inputBlocks {
		u, err := setter(block)
		if err != nil {
			return err
		}
		resultBlocks = append(resultBlocks, u)
	}
	bn.blocks = resultBlocks
	return nil
}

func (bn *BigNumber) GetHex() (hex string) {
	return bn.getValue(16, func(u Uint) string {
		return u.GetHex()
	})
}

func (bn *BigNumber) GetBinary() (hex string) {
	return bn.getValue(64, func(u Uint) string {
		return u.GetBinary()
	})
}

func (bn *BigNumber) getValue(blockSize int, getter func(Uint) string) (result string) {
	for i, block := range bn.blocks {
		blockValue := getter(block)
		if i != len(bn.blocks)-1 {
			missingZerosCount := blockSize - len(blockValue)
			var sb strings.Builder
			for i := 0; i < missingZerosCount; i++ {
				sb.WriteString("0")
			}
			sb.WriteString(blockValue)
			blockValue = sb.String()
		}
		result = blockValue + result
	}
	return
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

func binaryOperation(a, b BigNumber, operation func(Uint, Uint) Uint) (result BigNumber) {
	blocks := make([]Uint, 0)
	for i := 0; i < len(a.blocks) || i < len(b.blocks); i++ {
		if i >= len(a.blocks) {
			blocks = append(blocks, operation(b.blocks[i], Uint{0}))
		} else if i >= len(b.blocks) {
			blocks = append(blocks, operation(a.blocks[i], Uint{0}))
		} else {
			blocks = append(blocks, operation(a.blocks[i], b.blocks[i]))
		}
	}
	result.blocks = blocks
	return
}

func (bn *BigNumber) XOR(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.XOR(u2) })
}

func (bn *BigNumber) AND(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.AND(u2) })
}

func (bn *BigNumber) OR(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.OR(u2) })
}

func (bn *BigNumber) ShiftL(n int) (result BigNumber) {
	binary := bn.GetBinary()
	var sb strings.Builder
	sb.WriteString(binary)
	for i := 0; i < n; i++ {
		sb.WriteRune('0')
	}
	result.SetBinary(sb.String())
	return
}

func (bn *BigNumber) ShiftR(n int) (result BigNumber) {
	binary := bn.GetBinary()
	result.SetBinary(binary[:len(binary)-n])
	return
}
