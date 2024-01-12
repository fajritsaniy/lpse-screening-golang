package utils

import "strings"

func ReplaceSpacesWithPlus(input string) string {
	// Replace spaces with "+"
	result := strings.Replace(input, " ", "+", -1)
	return result
}
