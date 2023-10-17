# Big Numbers
Cryptography for developers by Distributed Lab - Practice #2

## Work done
* `src/bignumbers.go` - the implementation of a Big Number.

* `src/uint.go` - the data type used by `BigNumber`. Just a convenience that lets you convert a `uint64` to different notations and perform basic binary and arithmetic operations.

* `arithmeticops.go/binaryops.go` - interfaces that the `BigNumber` needs to implement. 

* `src/util` - utility functions.

* `src/validators` - set of functions to validate values.

* `tests/` - contains the tests.

## List of operations
| **#** 	| **Operation** 	| **Name** 	|
|---	|---	|---	|
| 1 	| INV 	| bitwise inversion 	|
| 2 	| XOR 	| bitwise exclusive or 	|
| 3 	| OR 	| bitwise or 	|
| 4 	| AND 	| bitwise and 	|
| 5 	| shiftR 	| shift to the right by n bits 	|
| 6 	| shiftL 	| shift to the left by n bits 	|
| 7 	| ADD 	| addition 	|
| 8 	| SUB 	| subtraction 	|
| 9 	| MOD 	| modulo 	|

## Testing

To run the tests, execute one of the following commands in the root folder:

```bash
go test .\tests\
```
```bash
go test -v .\tests\
```

When pushing, all the tests run automatically with the GitHub Actions.

## Example

```go
import (
	"fmt"

	bignumbers "github.com/danielost/big-numbers/src"
)

func foo() {
	var bnA bignumbers.BigNumber
	var bnB bignumbers.BigNumber

	bnA.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	bnB.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")

	bnC := bnA.XOR(bnB)
	fmt.Println(bnC.GetHex())
}
```
