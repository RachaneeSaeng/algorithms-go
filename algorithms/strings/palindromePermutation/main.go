package main

// Check if a given string is a permutation of a palindrome
// palindrome is a string that is the same forwards and backwords e.g. abcba
func isPermutationOfPalindrome(phase string) bool {
	if len(phase) == 0 {
		return false
	}

	bitVector := buildBitVector(phase)
	return checkOnlyOneBitSet(bitVector)
}

func buildBitVector(phase string) int32 {
	var bitVector int32
	for _, c := range phase {
		index := getCharNumber(c)
		if index >= 0 {
			bitVector = toggleBit(bitVector, index)
		}
	}

	return bitVector
}

func getCharNumber(char rune) int {
	if 'a' <= char && char <= 'z' {
		return int(char - 'a')
	}
	return -1
}

func toggleBit(bitVector int32, index int) int32 {
	return (bitVector ^ (1 << uint(index)))
}

func checkOnlyOneBitSet(bitVector int32) bool {
	return (bitVector & (bitVector - 1)) == 0
}

func main() {

}
