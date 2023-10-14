package bignumbers

import (
	"math"
)

func reverseString(s string) string {
	runeArray := []rune(s)
	for i := 0; i < len(runeArray)/2; i++ {
		runeArray[i], runeArray[len(runeArray)-i-1] = runeArray[len(runeArray)-i-1], runeArray[i]
	}
	return string(runeArray)
}

func breakStringIntoBlocks(input string, blockSize int) []string {
	numberOfBlocks := uint(math.Ceil(float64(len(input)) / float64(blockSize)))
	blocks := make([]string, numberOfBlocks)
	reversedInput := reverseString(input)

	for i := 0; i < int(numberOfBlocks); i++ {
		if i*blockSize+blockSize > len(reversedInput) {
			blocks[i] = reverseString(reversedInput[i*blockSize:])
		} else {
			blocks[i] = reverseString(reversedInput[i*blockSize : i*blockSize+blockSize])
		}
	}

	return blocks
}
