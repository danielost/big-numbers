package bignumbers

type ArithmeticOps interface {
	ADD(interface{})
	SUB(interface{})
	MOD(interface{})
}

type AdvancedArithmeticOps interface {
	MUL(interface{})
	DIV(interface{})
	POWMOD(interface{}, int)
}
