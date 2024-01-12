package utils

import (
	"regexp"
	"strings"
)

func RemoveExtraSpaces(input string) string {
	// Replace multiple consecutive spaces with a single space
	re := regexp.MustCompile(`\s+`)
	noExtraSpaces := re.ReplaceAllString(input, " ")

	// Trim leading and trailing spaces
	trimmed := strings.TrimSpace(noExtraSpaces)

	return trimmed
}
