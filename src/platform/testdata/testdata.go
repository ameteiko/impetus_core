package testdata

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v5"
)

const (
	randomnessLength = 6
)

// Slug returns a random slug with an optional prefix.
func Slug(prefix string) string {
	return gofakeit.Numerify(fmt.Sprintf("%s%s", prefix, strings.Repeat("#", randomnessLength)))
}
