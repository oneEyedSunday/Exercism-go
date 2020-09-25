// Package proverb makes proverbs from input strings
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) > 0 {
		outputs := make([]string, len(rhyme))

		if len(rhyme) > 1 {
			for index, item := range rhyme[:len(rhyme)-1] {
				outputs[index] = fmt.Sprintf("For want of a %s the %s was lost.", item, rhyme[index+1])
			}
		}

		outputs[len(outputs)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		return outputs
	}
	return []string{}
}
