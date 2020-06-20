package sanitiser

import (
	"strings"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

const (
	whitespaceReplacement = "-"
)

var kindMappings = map[string][]string{
	// TODO: Refactor this dependency for the value-objects package.
	"artefact": {"artifact", "artfct", "artf", "art", "atf"},
}

type Node struct{}

func (s Node) Sanitise(raw dto.Node) (node dto.Node) {
	node = raw
	node.Content = trim(raw.Content)
	node.ContentURI = trim(raw.ContentURI)
	node.Kind = s.sanitiseKind(raw.Kind)
	node.Slug = s.sanitiseSlug(raw.Slug, raw.Title)
	node.Title = s.sanitiseTitle(raw.Title)

	return node
}

// sanitiseKind performs:
//  - removing of starting and trailing whitespaces
//  - mapping of the shortcuts and misspelled values to a correct one.
//
// TODO: maybe return a default value Artifact?
func (_ Node) sanitiseKind(k string) string {
	k = trim(k)

	for kind, shortcuts := range kindMappings {
		if kind == k {
			return kind
		}

		for _, shortcut := range shortcuts {
			if shortcut == k {
				return kind
			}
		}
	}

	return k
}

// sanitiseSlug performs slug sanitising. If slug parameter is empty, then the title is used to generate a slug.
// sanitiseSlug performs:
// 	- cleaning for starting and trailing whitespaces
//  - casts the value to the lover case
//  - replaces all the spaces to the dashes.
func (_ Node) sanitiseSlug(s string, t string) string {
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
func (_ Node) sanitiseTitle(t string) string {
	t = trim(t)
	t = mergeWhitespaces(t)

	return t
}

// sanitiseTags performs:
//  - trimming of starting and trailing whitespaces
//  - replaces several sequential whitespaces with a single one
//  - removes duplicated tags.
func (_ Node) sanitiseTags(tags []string) (sanitisedTags []string) {
	if tags == nil || len(tags) == 0 {
		return nil
	}

	tagsMap := make(map[string]struct{})
	for _, tag := range tags {
		tag = mergeWhitespaces(trim(tag))
		if _, exists := tagsMap[tag]; exists {
			continue
		}

		tagsMap[tag] = struct{}{}
		sanitisedTags = append(sanitisedTags, tag)
	}

	return sanitisedTags
}
