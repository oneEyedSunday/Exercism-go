package hamming

import (
	"errors"
)

// Distance calculates the hamming distance or returns an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String lengths do not match")
	}

	mismatches := 0
	for index := range a {
		if a[index] != b[index] {
			mismatches++
		}
	}

	return mismatches, nil
}
