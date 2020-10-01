package raindrops

import (
	"strconv"
)

// Convert transforms a number into a string that contains raindrop sounds
func Convert(value int) string {
	raindrop := ""
	if value%3 == 0 {
		raindrop += "Pling"
	}

	if value%5 == 0 {
		raindrop += "Plang"
	}

	if value%7 == 0 {
		raindrop += "Plong"
	}

	if raindrop == "" {
		return strconv.Itoa(value)
	}

	return raindrop
}
