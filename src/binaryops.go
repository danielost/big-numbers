package bignumbers

type BinaryOps interface {
	Invert() BigNumber
	XOR(BigNumber) BigNumber
	OR(BigNumber) BigNumber
	AND(BigNumber) BigNumber
	ShiftR(int) BigNumber
	ShiftL(int) BigNumber
}
