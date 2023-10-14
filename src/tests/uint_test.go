package bignumbers_test

import (
	"math"
	"strings"
	"testing"

	bignumbers "github.com/danielost/big-numbers/src"
)

func TestUint_GetHex(t *testing.T) {
	var bn1 bignumbers.Uint
	var bn2 bignumbers.Uint
	bn1.SetDecimal(81985529216486895)
	bn2.SetDecimal(math.MaxUint64)
	tests := []struct {
		name        string
		u           *bignumbers.Uint
		expectedHex string
	}{
		{name: "Get 81985529216486895", u: &bn1, expectedHex: "123456789aBcDeF"},
		{name: "Get math.MaxUint64", u: &bn2, expectedHex: "FFFFFFFFFFFFFFFF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if hex := tt.u.GetHex(); hex != strings.ToLower(tt.expectedHex) {
				t.Errorf("Uint.GetHex() error: expected %s but got %s", tt.expectedHex, hex)
			}
		})
	}
}

func TestUint_SetHex(t *testing.T) {
	tests := []struct {
		name    string
		u       *bignumbers.Uint
		hex     string
		wantErr bool
	}{
		{name: "Set 123456789aBcDeF", u: new(bignumbers.Uint), hex: "123456789aBcDeF", wantErr: false},
		{name: "Set FFFFFFFFFFFFFFFF", u: new(bignumbers.Uint), hex: "FFFFFFFFFFFFFFFF", wantErr: false},
		{name: "Set abcdef12541abcdef2: length > 16", u: new(bignumbers.Uint), hex: "abcdef12541abcdef2", wantErr: true},
		{name: "Set eF3X7: contains prohibited symbol", u: new(bignumbers.Uint), hex: "eF3X7", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SetHex(tt.hex); (err != nil) != tt.wantErr {
				t.Errorf("Uint.SetHex() error = %v, wantErr %v", err, tt.wantErr)
			}
			if outputHex := tt.u.GetHex(); (outputHex != strings.ToLower(tt.hex)) != tt.wantErr {
				t.Errorf("Uint.SetHex() error: expected %s but got %s, wantErr %v", tt.hex, outputHex, tt.wantErr)
			}
		})
	}
}

func TestUint_GetBinary(t *testing.T) {
	var bn1 bignumbers.Uint
	var bn2 bignumbers.Uint
	bn1.SetDecimal(81985529216486895)
	bn2.SetDecimal(math.MaxUint64)
	tests := []struct {
		name           string
		u              *bignumbers.Uint
		expectedBinary string
	}{
		{name: "Get 81985529216486895", u: &bn1, expectedBinary: "100100011010001010110011110001001101010111100110111101111"},
		{name: "Get math.MaxUint64", u: &bn2, expectedBinary: "1111111111111111111111111111111111111111111111111111111111111111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if binary := tt.u.GetBinary(); binary != tt.expectedBinary {
				t.Errorf("Uint.GetBinary() error: expected %s but got %s", tt.expectedBinary, binary)
			}
		})
	}
}

func TestUint_SetBinary(t *testing.T) {
	tests := []struct {
		name    string
		u       *bignumbers.Uint
		binary  string
		wantErr bool
	}{
		{name: "Set 100100011010001010110011110001001101010111100110111101111", u: new(bignumbers.Uint), binary: "100100011010001010110011110001001101010111100110111101111", wantErr: false},
		{name: "Set 1111111111111111111111111111111111111111111111111111111111111111", u: new(bignumbers.Uint), binary: "1111111111111111111111111111111111111111111111111111111111111111", wantErr: false},
		{name: "Set 111111111110011111111111111111111111111111111101111111111111111010101010: length > 16", u: new(bignumbers.Uint), binary: "111111111110011111111111111111111111111111111101111111111111111010101010", wantErr: true},
		{name: "Set 10101201: contains prohibited symbol", u: new(bignumbers.Uint), binary: "10101201", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.SetBinary(tt.binary); (err != nil) != tt.wantErr {
				t.Errorf("Uint.SetBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
			if outputBinary := tt.u.GetBinary(); (outputBinary != tt.binary) != tt.wantErr {
				t.Errorf("Uint.SetBinary() error: expected %s but got %s, wantErr %v", tt.binary, outputBinary, tt.wantErr)
			}
		})
	}
}
