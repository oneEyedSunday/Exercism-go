// Package leap solves for leap years
package leap

func evenlyDivisibleBy(item int, value int) bool {
	return item%value == 0
}

// IsLeapYear determines if a year is leap or otherwise.
func IsLeapYear(year int) bool {
	return evenlyDivisibleBy(year, 4) && (!evenlyDivisibleBy(year, 100) || evenlyDivisibleBy(year, 400))
}
