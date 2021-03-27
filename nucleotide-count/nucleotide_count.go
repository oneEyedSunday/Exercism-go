package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune] int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[rune]int {'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, strand := range d {
		if _, found := h[strand]; !found {
			return make(Histogram), errors.New("invalid strand")
		}

		h[strand] += 1
	}
	return h, nil
}
