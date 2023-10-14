package bignumbers

type BinaryOps interface {
	Invert()
	XOR(BigNumber)
	OR(BigNumber)
	AND(BigNumber)
	ShiftR(int)
	ShiftL(int)
}
