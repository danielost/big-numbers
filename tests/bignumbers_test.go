package bignumbers_test

import (
	"strings"
	"testing"

	bignumbers "github.com/danielost/big-numbers/src"
)

func testEquality(a, b []bignumbers.Uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBigNumber_SetHex(t *testing.T) {
	tests := []struct {
		name           string
		hex            string
		expectedBlocks []bignumbers.Uint
	}{
		{name: "Set #1", hex: "1abc0000000dddddddddddddd0000ffffffff003", expectedBlocks: []bignumbers.Uint{{Value: 14987997152075051011}, {Value: 3903119677054429}, {Value: 448528384}}},
		{name: "Set #2", hex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", expectedBlocks: []bignumbers.Uint{{Value: 3347068819741802172}, {Value: 13920789932245924882}, {Value: 5626733489596141559}, {Value: 3733152895074749161}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bn bignumbers.BigNumber
			bn.SetHex(tt.hex)
			if !testEquality(bn.GetBlocks(), tt.expectedBlocks) {
				t.Errorf("BigNumber.SetHex() error: expected %v but got %v", tt.expectedBlocks, bn.GetBlocks())
			}
		})
	}
}

func TestBigNumber_GetHex(t *testing.T) {
	tests := []struct {
		name        string
		expectedHex string
	}{
		{name: "Get #1", expectedHex: "1abc0000000dddddddddddddd0000ffffffff003"},
		{name: "Get #2", expectedHex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bn bignumbers.BigNumber
			bn.SetHex(tt.expectedHex)
			if hex := bn.GetHex(); hex != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.GetHex() error: expected %s but got %s", tt.expectedHex, hex)
			}
		})
	}
}

func TestBigNumber_SetBinary(t *testing.T) {
	tests := []struct {
		name           string
		binary         string
		expectedBlocks []bignumbers.Uint
	}{
		{name: "SetBinary #1", binary: "1010101111001101111011110000000100100011010001010110011110001001111111101101110010111010100110000111011001010100001100100001000010101011110011011110111100000001001000110100010101100111100010011111111011011100101110101001100001110110010101000011001000010000", expectedBlocks: []bignumbers.Uint{{Value: 18364758544493064720}, {Value: 12379813738877118345}, {Value: 18364758544493064720}, {Value: 12379813738877118345}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bn bignumbers.BigNumber
			bn.SetBinary(tt.binary)
			if !testEquality(bn.GetBlocks(), tt.expectedBlocks) {
				t.Errorf("BigNumber.SetBinary() error: expected %v but got %v", tt.expectedBlocks, bn.GetBlocks())
			}
		})
	}
}

func TestBigNumber_GetBinary(t *testing.T) {
	tests := []struct {
		name           string
		hex            string
		expectedBinary string
	}{
		{name: "GetBinary #1", hex: "1abc0000000dddddddddddddd0000ffffffff003", expectedBinary: "1101010111100000000000000000000000000000011011101110111011101110111011101110111011101110111011101000000000000000011111111111111111111111111111111000000000011"},
		{name: "GetBinary #2", hex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", expectedBinary: "11001111001110110100101100011101101011001001101100101011101001010011100001011000101100010011000000110100101100000011111111011111000001001100001001010010110000000110000101101000111100000100100010111001110011001011010101101110100111011111101111111010111100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bn bignumbers.BigNumber
			bn.SetHex(tt.hex)
			if hex := bn.GetBinary(); hex != strings.ToLower(tt.expectedBinary) {
				t.Errorf("BigNumber.GetBinary() error: expected %s but got %s", tt.expectedBinary, hex)
			}
		})
	}
}

func TestBigNumber_Invert(t *testing.T) {
	tests := []struct {
		name        string
		hex         string
		expectedHex string
	}{
		{name: "Invert #1", hex: "1abc0000000dddddddddddddd0000ffffffff003", expectedHex: "E543FFFFFFF22222222222222FFFF00000000FFC"},
		{name: "Invert #2", hex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", expectedHex: "cc312d3894d93516b1e9d3b3f2d3f0083ecf6b4fe7a5c3edd18cd2a458810143"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bn bignumbers.BigNumber
			bn.SetHex(tt.hex)
			if bnInverted := bn.Invert(); bnInverted.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.Invert() error: expected %s but got %s", tt.expectedHex, bnInverted.GetHex())
			}
		})
	}
}

func TestBigNumber_XOR(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
	}{
		{name: "XOR #1", left: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", right: "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", expectedHex: "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8"},
		{name: "XOR #2", left: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", right: "0", expectedHex: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"},
	}
	for _, tt := range tests {
		var bnLeft bignumbers.BigNumber
		var bnRight bignumbers.BigNumber
		bnLeft.SetHex(tt.left)
		bnRight.SetHex(tt.right)
		t.Run(tt.name, func(t *testing.T) {
			if result := bnLeft.XOR(bnRight); result.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.XOR() error: expected %s but got %s", tt.expectedHex, result.GetHex())
			}
		})
	}
}

func TestBigNumber_AND(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
	}{
		{name: "AND #1", left: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", right: "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", expectedHex: "403d208400a113220340808088d16a1b10121078400c1002748196dd62460204"},
		{name: "AND #2", left: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", right: "0", expectedHex: ""},
	}
	for _, tt := range tests {
		var bnLeft bignumbers.BigNumber
		var bnRight bignumbers.BigNumber
		bnLeft.SetHex(tt.left)
		bnRight.SetHex(tt.right)
		t.Run(tt.name, func(t *testing.T) {
			if result := bnLeft.AND(bnRight); result.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.AND() error: expected %s but got %s", tt.expectedHex, result.GetHex())
			}
		})
	}
}

func TestBigNumber_OR(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
	}{
		{name: "OR #1", left: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", right: "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c", expectedHex: "51bff8ad9cafd72eabffbfc9befffffffcffbffaffdd779afdf3d7fdf7f73fbc"},
		{name: "OR #2", left: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", right: "0", expectedHex: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"},
	}
	for _, tt := range tests {
		var bnLeft bignumbers.BigNumber
		var bnRight bignumbers.BigNumber
		bnLeft.SetHex(tt.left)
		bnRight.SetHex(tt.right)
		t.Run(tt.name, func(t *testing.T) {
			if result := bnLeft.OR(bnRight); result.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.OR() error: expected %s but got %s", tt.expectedHex, result.GetHex())
			}
		})
	}
}

func TestBigNumber_ShiftL(t *testing.T) {
	tests := []struct {
		name        string
		hex         string
		shiftBy     int
		expectedHex string
	}{
		{name: "LeftShift #1", hex: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", shiftBy: 3, expectedHex: "28DFB0420A56AB9351E0DF604C7BBD8DAA7FD93C3FC6A9453A60EBFEF32387520"},
		{name: "LeftShift #2", hex: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", shiftBy: 64, expectedHex: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF0000000000000000"},
	}
	for _, tt := range tests {
		var bn bignumbers.BigNumber
		bn.SetHex(tt.hex)
		t.Run(tt.name, func(t *testing.T) {
			if result := bn.ShiftL(tt.shiftBy); result.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.ShiftL() error: expected %s but got %s", tt.expectedHex, result.GetHex())
			}
		})
	}
}

func TestBigNumber_ShiftR(t *testing.T) {
	tests := []struct {
		name        string
		hex         string
		shiftBy     int
		expectedHex string
	}{
		{name: "RightShift #1", hex: "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4", shiftBy: 3, expectedHex: "A37EC108295AAE4D47837D8131EEF636A9FF64F0FF1AA514E983AFFBCC8E1D4"},
		{name: "RightShift #2", hex: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", shiftBy: 10, expectedHex: "3FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"},
	}
	for _, tt := range tests {
		var bn bignumbers.BigNumber
		bn.SetHex(tt.hex)
		t.Run(tt.name, func(t *testing.T) {
			if result := bn.ShiftR(tt.shiftBy); result.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.ShiftR() error: expected %s but got %s", tt.expectedHex, result.GetHex())
			}
		})
	}
}

func TestBigNumber_ADD(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
	}{
		{name: "ADD #1", left: "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80", right: "70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb", expectedHex: "a78865c13b14ae4e25e90771b54963ee2d68c0a64d4a8ba7c6f45ee0e9daa65b"},
		{name: "ADD #2", left: "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80", right: "be024579b3f1357908641fdb97ffffffacf13468ad44444497530eca86ffffff", expectedHex: "F4F26DD1BFA162412F8EB9DDA74200E2F3D3AB1713928A3317C7643F69F5AB7F"},
		{name: "ADD #3", left: "10", right: "20", expectedHex: "30"},
		{name: "ADD #4", left: "FFFFFFFFFFFFFFFF", right: "1", expectedHex: "10000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bnLeft bignumbers.BigNumber
			var bnRight bignumbers.BigNumber
			bnLeft.SetHex(tt.left)
			bnRight.SetHex(tt.right)
			if sum := bnLeft.ADD(bnRight); sum.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.ADD() error: expected %s but got %s", tt.expectedHex, sum.GetHex())
			}
		})
	}
}

func TestBigNumber_SUB(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
		wantErr     bool
	}{
		{name: "SUB #1", left: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", right: "22e962951cb6cd2ce279ab0e2095825c141d48ef3ca9dabf253e38760b57fe03", expectedHex: "10e570324e6ffdbc6b9c813dec968d9bad134bc0dbb061530934f4e59c2700b9", wantErr: false},
		{name: "SUB #2", left: "10abcdef0123456789fedcba9876543210", right: "b97ffffffacf13468ad44444497530eca86ffffff", expectedHex: "", wantErr: true},
		{name: "SUB #3", left: "abcdef0123456789fedcba9876543210abcdef0123456789fedcba9876543210", right: "1234567890abcdef0987654321abcdef0123456789fedcba9876543210abcdef", expectedHex: "999998889299999AF555555554A86421AAAAA99999468ACF6666666665A86421", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bnLeft bignumbers.BigNumber
			var bnRight bignumbers.BigNumber
			bnLeft.SetHex(tt.left)
			bnRight.SetHex(tt.right)
			diff, err := bnLeft.SUB(bnRight)
			if err != nil && !tt.wantErr {
				t.Errorf("BigNumber.SUB() error: %v", err)
			}
			if diff.GetHex() != strings.ToLower(tt.expectedHex) && !tt.wantErr {
				t.Errorf("BigNumber.SUB() error: expected %s but got %s", tt.expectedHex, diff.GetHex())
			}
		})
	}
}

func TestBigNumber_MOD(t *testing.T) {
	tests := []struct {
		name        string
		left        string
		right       string
		expectedHex string
	}{
		{name: "MOD #1", left: "ABCDEF", right: "123456", expectedHex: "7f6e9"},
		{name: "MOD #2", left: "9876543", right: "123456789", expectedHex: "9876543"},
		{name: "MOD #3", left: "abcdef0123456789fedcba9876543210abcdef0123456789fedcba9876543210", right: "1234567890abcdef0987654321abcdef0123456789fedcba9876543210abcdef", expectedHex: "7f6e4c40d3b2a22a91a2b3c4749f4a9a1907e5d494fa4faa2b3c4d5e049f4a9"},
		{name: "MOD #4", left: "123456", right: "1", expectedHex: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bnLeft bignumbers.BigNumber
			var bnRight bignumbers.BigNumber
			bnLeft.SetHex(tt.left)
			bnRight.SetHex(tt.right)
			if sum := bnLeft.MOD(bnRight); sum.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.ADD() error: expected %s but got %s", tt.expectedHex, sum.GetHex())
			}
		})
	}
}
