package strand

// ToRNA transcribes strings 
func ToRNA(dna string) string {
	output := make([] rune, len(dna))

	lookup := map[rune]rune {
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for index, tna := range dna {
		output[index] = lookup[tna]
	}

	return string(output)
}
