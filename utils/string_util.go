package utils

import "strings"

func IsBlank(val string) bool {
	trimmedVal := strings.TrimSpace(val)

	return trimmedVal == ""
}

func IsNotBlank(val string) bool {
	return !IsBlank(val)
}
