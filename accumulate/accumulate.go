package accumulate

// Accumulate applies a function (operation) on a collection, returning a new collection
func Accumulate(input []string, op func(string) string) []string {
	var output []string = make([]string, len(input))

	for index, item := range input {
		output[index] = op(item)
	}
	return output
}
