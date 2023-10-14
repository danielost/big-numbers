package bignumbers

type ArithmeticOps interface {
	ADD(BigNumber)
	SUB(BigNumber)
	MOD(BigNumber)
}

type AdvancedArithmeticOps interface {
	MUL(BigNumber)
	DIV(BigNumber)
	POWMOD(BigNumber, int)
}
