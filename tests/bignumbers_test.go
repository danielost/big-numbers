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
		bn             *bignumbers.BigNumber
		hex            string
		expectedBlocks []bignumbers.Uint
	}{
		{name: "Set 1abc0000000dddddddddddddd0000ffffffff003", bn: new(bignumbers.BigNumber), hex: "1abc0000000dddddddddddddd0000ffffffff003", expectedBlocks: []bignumbers.Uint{{Value: 14987997152075051011}, {Value: 3903119677054429}, {Value: 448528384}}},
		{name: "Set 33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", bn: new(bignumbers.BigNumber), hex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", expectedBlocks: []bignumbers.Uint{{Value: 3347068819741802172}, {Value: 13920789932245924882}, {Value: 5626733489596141559}, {Value: 3733152895074749161}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bn.SetHex(tt.hex)
			if !testEquality(tt.bn.GetBlocks(), tt.expectedBlocks) {
				t.Errorf("BigNumber.SetHex() error: expected %v but got %v", tt.expectedBlocks, tt.bn.GetBlocks())
			}
		})
	}
}

func TestBigNumber_GetHex(t *testing.T) {
	var bn1 bignumbers.BigNumber
	var bn2 bignumbers.BigNumber
	bn1.SetHex("1abc0000000dddddddddddddd0000ffffffff003")
	bn2.SetHex("33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc")
	tests := []struct {
		name        string
		bn          *bignumbers.BigNumber
		expectedHex string
	}{
		{name: "Get 1abc0000000dddddddddddddd0000ffffffff003", bn: &bn1, expectedHex: "1abc0000000dddddddddddddd0000ffffffff003"},
		{name: "Get 33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", bn: &bn2, expectedHex: "33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if hex := tt.bn.GetHex(); hex != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.GetHex() error: expected %s but got %s", tt.expectedHex, hex)
			}
		})
	}
}

func TestBigNumber_Invert(t *testing.T) {
	var bn1 bignumbers.BigNumber
	var bn2 bignumbers.BigNumber
	bn1.SetHex("1abc0000000dddddddddddddd0000ffffffff003")
	bn2.SetHex("33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc")
	tests := []struct {
		name        string
		bn          *bignumbers.BigNumber
		expectedHex string
	}{
		{name: "Invert 1abc0000000dddddddddddddd0000ffffffff003", bn: &bn1, expectedHex: "E543FFFFFFF22222222222222FFFF00000000FFC"},
		{name: "Invert 33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc", bn: &bn2, expectedHex: "cc312d3894d93516b1e9d3b3f2d3f0083ecf6b4fe7a5c3edd18cd2a458810143"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if bnInverted := tt.bn.Invert(); bnInverted.GetHex() != strings.ToLower(tt.expectedHex) {
				t.Errorf("BigNumber.Invert() error: expected %s but got %s", tt.expectedHex, bnInverted.GetHex())
			}
		})
	}
}

func TestBigNumber_XOR(t *testing.T) {
	var bn1 bignumbers.BigNumber
	var bn2 bignumbers.BigNumber
	bn1.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	bn2.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	expectedHex := "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8"
	t.Run("XOR", func(t *testing.T) {
		if result := bn1.XOR(bn2); result.GetHex() != strings.ToLower(expectedHex) {
			t.Errorf("BigNumber.XOR() error: expected %s but got %s", expectedHex, result.GetHex())
		}
	})
}

func TestBigNumber_AND(t *testing.T) {
	var bn1 bignumbers.BigNumber
	var bn2 bignumbers.BigNumber
	bn1.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	bn2.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	expectedHex := "403d208400a113220340808088d16a1b10121078400c1002748196dd62460204"
	t.Run("AND", func(t *testing.T) {
		if result := bn1.AND(bn2); result.GetHex() != strings.ToLower(expectedHex) {
			t.Errorf("BigNumber.AND() error: expected %s but got %s", expectedHex, result.GetHex())
		}
	})
}

func TestBigNumber_OR(t *testing.T) {
	var bn1 bignumbers.BigNumber
	var bn2 bignumbers.BigNumber
	bn1.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	bn2.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	expectedHex := "51bff8ad9cafd72eabffbfc9befffffffcffbffaffdd779afdf3d7fdf7f73fbc"
	t.Run("OR", func(t *testing.T) {
		if result := bn1.OR(bn2); result.GetHex() != strings.ToLower(expectedHex) {
			t.Errorf("BigNumber.OR() error: expected %s but got %s", expectedHex, result.GetHex())
		}
	})
}
