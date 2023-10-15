package bignumbers

type BinaryOps interface {
	Invert()
	XOR(interface{})
	OR(interface{})
	AND(interface{})
	ShiftR(int)
	ShiftL(int)
}
