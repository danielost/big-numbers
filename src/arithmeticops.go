package bignumbers

type ArithmeticOps interface {
	ADD(BigNumber) BigNumber
	SUB(BigNumber) BigNumber
	MOD(BigNumber) BigNumber
}

type AdvancedArithmeticOps interface {
	MUL(BigNumber) BigNumber
	DIV(BigNumber) BigNumber
	POWMOD(BigNumber, uint64) BigNumber
}
