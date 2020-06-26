package sanitiser

import (
	"strings"
)

const (
	whitespaceReplacement = "-"
)

type Sanitiser struct {}

func NewSanitiser() Sanitiser {
	return Sanitiser{}
}

func (_ Sanitiser) SanitiseID(id string) string {
	id = trim(id)

	return trim(id)
}


// sanitiseSlug performs slug sanitising. If slug parameter is empty, then the title is used to generate a slug.
// sanitiseSlug performs:
// 	- cleaning for starting and trailing whitespaces
//  - casts the value to the lover case
//  - replaces all the spaces to the dashes.
func (_ Sanitiser) SanitiseSlug(s string, t string) string {
	s = trim(s)

	// Take the title as a default slug value.
	if s == "" {
		s = trim(t)
	}

	s = strings.ToLower(s)
	s = mergeWhitespaces(s)
	s = strings.ReplaceAll(s, " ", whitespaceReplacement)

	return s
}

// sanitiseTitle performs:
//   - cleaning for starting and trailing whitespaces
//   - replaces several spaces with one.
//
// TODO: Think of capitalising the words in the title: "Router for an http server" -> "Router for an HTTP Server".
func (_ Sanitiser) SanitiseTitle(t string) string {
	t = trim(t)
	t = mergeWhitespaces(t)

	return t
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func mergeWhitespaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
