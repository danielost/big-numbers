package bignumbers_test

import (
	"testing"

	bignumbers "github.com/danielost/big-numbers/src"
)

func TestValidateHex(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		wantErr bool
	}{
		{name: "Validate 123456789aBcDeF", hex: "123456789aBcDeF", wantErr: false},
		{name: "Validate abcdef12541abcdef2", hex: "abcdef12541abcdef2", wantErr: true},
		{name: "Validate eF3X7", hex: "eF3X7", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := bignumbers.ValidateHex(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestValidateBinary(t *testing.T) {
	tests := []struct {
		name    string
		binary  string
		wantErr bool
	}{
		{name: "Validate 100100011010001010110011110001001101010111100110111101111", binary: "100100011010001010110011110001001101010111100110111101111", wantErr: false},
		{name: "Validate 111111111110011111111111111111111111111111111101111111111111111010101010", binary: "111111111110011111111111111111111111111111111101111111111111111010101010", wantErr: true},
		{name: "Validate 10101201", binary: "10101201", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := bignumbers.ValidateBinary(tt.binary)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
