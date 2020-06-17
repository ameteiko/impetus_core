package sanitiser

import (
	"strings"
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func mergeWhitespaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
