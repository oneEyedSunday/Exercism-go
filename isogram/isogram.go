package isogram

import (
	"strings"
)

type Dict map[rune]bool

func IsIsogram (c string) bool {
	dict := make(Dict)
	whitelist := Dict { ' ': true, '-': true }
	for _, char := range strings.ToLower(c) {
		if !whitelist[char] && dict[char] == true {
			return false;
		}
		dict[char] = true;
	}
	return true
}
