package bignumbers

type BigNumber struct {
	blocks []Uint
}

func (bn *BigNumber) SetHex(hex string) error {
	stringBlocks := breakStringIntoBlocks(hex, 16)
	uintBlocks := make([]Uint, 0)
	for _, b := range stringBlocks {
		var u Uint
		if err := u.SetHex(b); err != nil {
			return err
		}
		uintBlocks = append(uintBlocks, u)
	}
	bn.blocks = uintBlocks

	return nil
}

func (bn *BigNumber) GetHex() (hex string) {
	for _, block := range bn.blocks {
		hex = block.GetHex() + hex
	}
	return hex
}
