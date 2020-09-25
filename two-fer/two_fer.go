// Package twofer holds my solution to the two_fer Exercism problem .
package twofer

import "fmt"

// ShareWith returns a string of the form "One for X, one for me.", where X is an input or 'you'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
