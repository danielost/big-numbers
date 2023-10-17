package bignumbers

import (
	"fmt"
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

func (bn *BigNumber) clearLeadingZeros() {
	for len(bn.blocks) > 0 && bn.blocks[len(bn.blocks)-1].GetDecimal() == 0 {
		bn.blocks = bn.blocks[:len(bn.blocks)-1]
	}
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

func (bn *BigNumber) LessThan(other BigNumber) bool {
	thisBlocks := bn.GetBlocks()
	otherBlocks := other.GetBlocks()
	if len(thisBlocks) < len(otherBlocks) {
		return true
	}
	if len(thisBlocks) > len(otherBlocks) {
		return false
	}
	for i := len(thisBlocks) - 1; i >= 0; i-- {
		if thisBlocks[i].GetDecimal() < otherBlocks[i].GetDecimal() {
			return true
		}
		if thisBlocks[i].GetDecimal() > otherBlocks[i].GetDecimal() {
			return false
		}
	}
	return false
}

func (bn *BigNumber) ADD(other BigNumber) (res BigNumber) {
	carry := Uint{0}
	thisBlocks := bn.blocks
	otherBlocks := other.blocks
	for i := 0; i < len(thisBlocks) || i < len(otherBlocks); i++ {
		if i >= len(thisBlocks) {
			res.blocks = append(res.blocks, otherBlocks[i].ADD(carry))
			carry = Uint{}
		} else if i >= len(otherBlocks) {
			res.blocks = append(res.blocks, thisBlocks[i].ADD(carry))
			carry = Uint{}
		} else {
			sum := thisBlocks[i].ADD(otherBlocks[i])
			sum = sum.ADD(carry)
			res.blocks = append(res.blocks, sum)
			if sum.GetDecimal() < thisBlocks[i].GetDecimal() || sum.GetDecimal() < otherBlocks[i].GetDecimal() {
				carry = Uint{1}
			} else {
				carry = Uint{0}
			}
		}
	}
	if carry.GetDecimal() > 0 {
		res.blocks = append(res.blocks, carry)
	}
	return
}

func (bn *BigNumber) SUB(other BigNumber) (res BigNumber, err error) {
	if bn.LessThan(other) {
		return BigNumber{}, fmt.Errorf("sub result is negative")
	}
	carry := Uint{0}
	thisBlocks := bn.blocks
	otherBlocks := other.blocks
	for i := 0; i < len(thisBlocks) || i < len(otherBlocks); i++ {
		if i >= len(thisBlocks) {
			res.blocks = append(res.blocks, Uint{otherBlocks[i].Value - carry.Value})
			carry = Uint{}
		} else if i >= len(otherBlocks) {
			res.blocks = append(res.blocks, Uint{thisBlocks[i].Value - carry.Value})
			carry = Uint{}
		} else {
			sum := thisBlocks[i].Value - otherBlocks[i].Value - carry.Value
			res.blocks = append(res.blocks, Uint{sum})
			if sum > thisBlocks[i].GetDecimal() {
				carry = Uint{1}
			} else {
				carry = Uint{0}
			}
		}
	}
	if carry.GetDecimal() > 0 {
		res.blocks = append(res.blocks, carry)
	}
	res.clearLeadingZeros()
	return
}
