package util

import (
	"strings"
)

func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func IsValidEmailString(str string) bool {

	if IsStringEmpty(str) {
		return false
	}

	return strings.Contains(str, "@")
}
