package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTitle(t *testing.T) {
	tcs := []struct {
		desc  string
		title string
		err   error
	}{
		{
			desc:  "ForAnEmptyTitle",
			title: "",
			err:   ErrTitleIsEmpty,
		},
		{
			desc:  "ForAnNotEmptyTitle",
			title: "some title",
			err:   nil,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			err := Validator{}.ValidateTitle(tc.title)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestValidateSlug(t *testing.T) {
	tcs := []struct {
		desc string
		slug string
		err  error
	}{
		{
			desc: "ForAnEmptyValue",
			slug: "",
			err:  ErrSlugIsEmpty,
		},
		{
			desc: "ForANotEmptyValue",
			slug: "some-slug",
			err:  nil,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			err := Validator{}.ValidateSlug(tc.slug)

			assert.Equal(t, tc.err, err)
		})
	}
}
