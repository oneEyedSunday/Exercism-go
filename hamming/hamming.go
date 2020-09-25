package hamming

import (
	"errors"
)

func getMismatched(first, second string) int {
	mismatches := 0
	for index := range first {
		if first[index] != second[index] {
			mismatches++
		}
	}

	return mismatches
}

// Distance calculates the hamming distance or returns an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String lengths do not match")
	}

	return getMismatched(a, b), nil
}
