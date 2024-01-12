package utils

import "fmt"

func InterfaceToString(data interface{}) string {
	// Use a type switch to check the type of data
	switch v := data.(type) {
	case int:
		// Convert int to string
		return fmt.Sprintf("%d", v)
	case float64:
		// Convert float64 to string
		return fmt.Sprintf("%f", v)
	case string:
		// If it's already a string, return as is
		return v
	default:
		// Handle other types accordingly or return an empty string
		return ""
	}
}
