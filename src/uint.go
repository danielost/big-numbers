package bignumbers

import (
	"fmt"
	"math"
	"strings"
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

func ValidateHex(hex string) (string, error) {
	hex = strings.ToLower(hex)
	if len(hex) > 16 {
		return "", fmt.Errorf("max hex length is 16")
	}
	for _, r := range hex {
		if !strings.ContainsRune(hexDigits, r) {
			return "", fmt.Errorf("'%s' is not a hex digit", string(r))
		}
	}
	return hex, nil
}

func ValidateBinary(bin string) error {
	if len(bin) > 64 {
		return fmt.Errorf("max binary length is 64")
	}
	for _, digit := range bin {
		if digit != '0' && digit != '1' {
			return fmt.Errorf("binary string must contain only ones and zeros")
		}
	}
	return nil
}
