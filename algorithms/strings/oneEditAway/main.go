package main

import (
	"math"
)

// Check if the second string is one (or zero) edit away from the first string
// edit means insert, replace, remove
func oneEditAway(first string, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}

	shortStr := first
	longStr := second
	if len(first) > len(second) {
		shortStr = second
		longStr = first
	}

	indexShort := 0
	indexLong := 0
	foundDiff := false
	for indexShort < len(shortStr) && indexLong < len(longStr) {
		if shortStr[indexShort] != longStr[indexLong] {
			if foundDiff {
				return false // more than one edit
			}

			foundDiff = true

			if len(shortStr) == len(longStr) { // replace
				indexShort++
				indexLong++
			} else { // insert case
				indexLong++ //Skip inserted char to check next char
			}
		} else {
			indexShort++
			indexLong++
		}
	}

	return true
}

func main() {

}
