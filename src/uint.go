package bignumbers

import (
	"fmt"
	"math"
)

type Uint struct {
	value uint64
}

const hexDigits string = "0123456789abcdef"

func (u *Uint) GetHex() (hex string) {
	value := u.value
	for value > 0 {
		hex = string(hexDigits[value%16]) + hex
		value /= 16
	}
	return
}

func (u *Uint) SetHex(hex string) error {
	validatedHex, err := ValidateHex(hex)
	if err != nil {
		return err
	}

	value := *new(uint64)
	for i, r := range validatedHex {
		pow := math.Pow(16, float64(len(validatedHex)-i-1))
		if r >= '0' && r <= '9' {
			left := uint64(r - '0')
			value += left * uint64(pow)
		} else {
			left := uint64(uint64(r-'a') + 10)
			value += left * uint64(pow)
		}
	}
	u.value = value

	return nil
}

func (u *Uint) GetDecimal() uint64 {
	return u.value
}

func (u *Uint) SetDecimal(value uint64) {
	u.value = value
}

func (u *Uint) GetBinary() (bin string) {
	value := u.value
	for value > 0 {
		bin = fmt.Sprint(value%2) + bin
		value /= 2
	}
	return
}

func (u *Uint) SetBinary(bin string) error {
	if err := ValidateBinary(bin); err != nil {
		return err
	}

	value := *new(uint64)
	for i, digit := range bin {
		left := uint64(digit - '0')
		right := uint64(math.Pow(2, float64(len(bin)-i-1)))
		value += left * right
	}
	u.value = value

	return nil
}

func (u *Uint) Invert() (result Uint) {
	result.SetDecimal(^u.value)
	return
}

func (u *Uint) XOR(other Uint) (result Uint) {
	result.SetDecimal(u.GetDecimal() ^ other.GetDecimal())
	return
}

func (u *Uint) AND(other Uint) (result Uint) {
	result.SetDecimal(u.GetDecimal() & other.GetDecimal())
	return
}

func (u *Uint) OR(other Uint) (result Uint) {
	result.SetDecimal(u.GetDecimal() | other.GetDecimal())
	return
}
