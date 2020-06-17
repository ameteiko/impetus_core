package sanitiser

import (
	"strings"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

const (
	whitespaceReplacement = "-"
)

type Node struct{}

func (s Node) Sanitise(raw dto.Node) (node dto.Node, err error) {
	node = raw
	node.Title = s.sanitiseTitle(raw.Title)
	node.Kind = s.sanitiseKind(raw.Kind)
	node.Slug = s.sanitiseSlug(raw.Slug)

	return node, nil
}

func (s Node) sanitiseKind(k string) string {
	return k
}

// sanitiseTitle performs:
//   - cleaning for starting and trailing whitespaces
//   - replaces several spaces with one.
func (_ Node) sanitiseTitle(t string) string {
	t = strings.TrimSpace(t)
	t = strings.Join(strings.Fields(t), " ")

	return t
}

// sanitiseSlug performs:
// 	- cleaning starting and trailing whitespaces
//  - casts the value to the lover case
//  - replaces all the spaces to the dashes.
func (_ Node) sanitiseSlug(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), whitespaceReplacement)

	return s
}
