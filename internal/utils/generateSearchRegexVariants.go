package utils

import (
	"strings"
)

// generateSearchTextVariants creates regex variants for fuzzy search.
func GenerateSearchTextVariants(query string, n int) string {
	var variants []string

	var replaceChars func(currentIndex int, indicesToReplace []int)
	replaceChars = func(currentIndex int, indicesToReplace []int) {
		if len(indicesToReplace) == n {
			runes := []rune(query)
			for _, idx := range indicesToReplace {
				runes[idx] = '.'
			}
			variants = append(variants, string(runes))
			return
		}
		for i := currentIndex; i < len(query); i++ {
			replaceChars(i+1, append(indicesToReplace, i))
		}
	}

	replaceChars(0, []int{})
	return strings.Join(variants, "|")
}
