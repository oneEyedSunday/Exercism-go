package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	var output map[string]int = make(map[string]int)

	for score, characters := range input {
		for _, character := range characters {
			output[strings.ToLower(character)] = score
		}
	}
	return output
}
