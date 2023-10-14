package bignumbers

import (
	"fmt"
	"strings"
)

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
