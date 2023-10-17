package bignumbers

import (
	"fmt"
	"strings"
)

// BigNumber represents a big number as a slice of blocks. Each block is a uint64 value.
type BigNumber struct {
	blocks []Uint
}

// GetBlocks returns the blocks of the BigNumber.
func (bn *BigNumber) GetBlocks() []Uint {
	return bn.blocks
}

// SetBlocks sets the blocks of the BigNumber to the provided slice of Uint blocks.
func (bn *BigNumber) SetBlocks(blocks []Uint) {
	bn.blocks = blocks
}

func (bn *BigNumber) AppendBlock(block Uint) {
	bn.SetBlocks(append(bn.GetBlocks(), block))
}

// SetHex sets the value of the BigNumber using a hexadecimal string.
func (bn *BigNumber) SetHex(hex string) error {
	return bn.setValue(16, hex, func(s string) (Uint, error) {
		var u Uint
		if err := u.SetHex(s); err != nil {
			return Uint{0}, err
		}
		return u, nil
	})
}

// SetBinary sets the value of the BigNumber using a binary string.
func (bn *BigNumber) SetBinary(binary string) error {
	return bn.setValue(64, binary, func(s string) (Uint, error) {
		var u Uint
		if err := u.SetBinary(s); err != nil {
			return Uint{0}, err
		}
		return u, nil
	})
}

// setValue sets the value of the BigNumber based on the provided block size, value, and setter function.
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
	bn.SetBlocks(resultBlocks)
	return nil
}

// GetHex returns the hexadecimal representation of the BigNumber.
func (bn *BigNumber) GetHex() (hex string) {
	return bn.getValue(16, func(u Uint) string {
		return u.GetHex()
	})
}

// GetBinary returns the binary representation of the BigNumber.
func (bn *BigNumber) GetBinary() (hex string) {
	return bn.getValue(64, func(u Uint) string {
		return u.GetBinary()
	})
}

// getValue returns the representation of the BigNumber based on the provided block size and getter function.
func (bn *BigNumber) getValue(blockSize int, getter func(Uint) string) (result string) {
	for i, block := range bn.GetBlocks() {
		blockValue := getter(block)
		if i != len(bn.GetBlocks())-1 {
			blockValue = AddLeadingZeros(blockValue, blockSize)
		}
		result = blockValue + result
	}
	return
}

// clearLeadingZeros removes leading zero blocks from the BigNumber.
func (bn *BigNumber) clearLeadingZeros() {
	for len(bn.GetBlocks()) > 0 && bn.GetBlocks()[len(bn.GetBlocks())-1].GetDecimal() == 0 {
		bn.SetBlocks(bn.GetBlocks()[:len(bn.GetBlocks())-1])
	}
}

// Invert returns the bitwise inversion of the BigNumber.
func (bn *BigNumber) Invert() (result BigNumber) {
	invertedBlocks := make([]Uint, len(bn.GetBlocks()))
	for i, block := range bn.GetBlocks() {
		invertedBlocks[i] = block.Invert()
	}
	result.SetBlocks(invertedBlocks)
	resultHex := result.GetHex()
	resultHex = resultHex[len(resultHex)-len(bn.GetHex()):]
	result.SetHex(resultHex)
	return
}

// binaryOperation performs a binary operation on two BigNumbers using the provided operation function.
func binaryOperation(a, b BigNumber, operation func(Uint, Uint) Uint) (result BigNumber) {
	aBlocks := a.GetBlocks()
	bBlocks := b.GetBlocks()
	for i := 0; i < len(aBlocks) || i < len(bBlocks); i++ {
		if i >= len(aBlocks) {
			result.AppendBlock(operation(bBlocks[i], Uint{0}))
		} else if i >= len(bBlocks) {
			result.AppendBlock(operation(aBlocks[i], Uint{0}))
		} else {
			result.AppendBlock(operation(aBlocks[i], bBlocks[i]))
		}
	}
	result.clearLeadingZeros()
	return
}

// XOR performs a bitwise XOR operation between two BigNumbers.
func (bn *BigNumber) XOR(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.XOR(u2) })
}

// AND performs a bitwise AND operation between two BigNumbers.
func (bn *BigNumber) AND(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.AND(u2) })
}

// OR performs a bitwise OR operation between two BigNumbers.
func (bn *BigNumber) OR(other BigNumber) (result BigNumber) {
	return binaryOperation(*bn, other, func(u1, u2 Uint) Uint { return u1.OR(u2) })
}

// ShiftL performs a left shift operation on the BigNumber.
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

// ShiftR performs a right shift operation on the BigNumber.
func (bn *BigNumber) ShiftR(n int) (result BigNumber) {
	binary := bn.GetBinary()
	result.SetBinary(binary[:len(binary)-n])
	return
}

// LessThan checks if the BigNumber is less than another BigNumber.
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

// ADD performs addition of two BigNumbers.
func (bn *BigNumber) ADD(other BigNumber) (result BigNumber) {
	carry := Uint{0}
	thisBlocks := bn.GetBlocks()
	otherBlocks := other.GetBlocks()
	for i := 0; i < len(thisBlocks) || i < len(otherBlocks); i++ {
		if i >= len(thisBlocks) {
			result.AppendBlock(otherBlocks[i].ADD(carry))
			carry = Uint{0}
		} else if i >= len(otherBlocks) {
			result.AppendBlock(thisBlocks[i].ADD(carry))
			carry = Uint{0}
		} else {
			sum := thisBlocks[i].ADD(otherBlocks[i])
			sum = sum.ADD(carry)
			result.AppendBlock(sum)
			if sum.GetDecimal() < thisBlocks[i].GetDecimal() || sum.GetDecimal() < otherBlocks[i].GetDecimal() {
				carry = Uint{1}
			} else {
				carry = Uint{0}
			}
		}
	}
	if carry.GetDecimal() > 0 {
		result.AppendBlock(carry)
	}
	return
}

// SUB performs subtraction of two BigNumbers.
func (bn *BigNumber) SUB(other BigNumber) (result BigNumber, err error) {
	if bn.LessThan(other) {
		return BigNumber{}, fmt.Errorf("sub result is negative")
	}
	carry := Uint{0}
	thisBlocks := bn.GetBlocks()
	otherBlocks := other.GetBlocks()
	for i := 0; i < len(thisBlocks) || i < len(otherBlocks); i++ {
		if i >= len(thisBlocks) {
			result.AppendBlock(otherBlocks[i].SUB(carry))
			carry = Uint{0}
		} else if i >= len(otherBlocks) {
			result.AppendBlock(thisBlocks[i].SUB(carry))
			carry = Uint{0}
		} else {
			diff := thisBlocks[i].SUB(otherBlocks[i])
			diff = diff.SUB(carry)
			result.AppendBlock(diff)
			if diff.GetDecimal() > thisBlocks[i].GetDecimal() {
				carry = Uint{1}
			} else {
				carry = Uint{0}
			}
		}
	}
	if carry.GetDecimal() > 0 {
		result.AppendBlock(carry)
	}
	result.clearLeadingZeros()
	return
}

/*
MOD calculates the modulo of two BigNumbers.
Note: this implementation is very inefficient as it doesn't use the division operation.
*/
func (bn *BigNumber) MOD(other BigNumber) (result BigNumber) {
	result.SetBlocks(bn.GetBlocks())
	for !result.LessThan(other) {
		result, _ = result.SUB(other)
	}
	return
}
