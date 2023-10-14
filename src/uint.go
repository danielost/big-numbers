package bignumbers

import (
	"fmt"
	"math"
	"strings"
)

type Uint uint64

const hexDigits string = "0123456789abcdef"

func (u Uint) ToHex() string {
	var sb strings.Builder

	for u > 0 {
		sb.WriteByte(hexDigits[u%16])
		u /= 16
	}

	return sb.String()
}

func FromHex(hex string) Uint {
	u := Uint(0)

	validatedHex, err := ValidateHex(hex)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	for i, r := range validatedHex {
		pow := math.Pow(16, float64(len(validatedHex)-i-1))
		if r >= 0 && r <= 9 {
			u += Uint(r-'0') * Uint(pow)
		} else {
			u += (Uint(r-'a') + 10) * Uint(pow)
		}
	}

	return u
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
