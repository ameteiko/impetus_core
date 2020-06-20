package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

func TestNodeValidate(t *testing.T) {
	tcs := []struct {
		desc  string
		title string
		slug  string
		errs  []error
	}{
		{
			desc:  "ForAnEmptySlugAndTitle",
			slug:  "",
			title: "",
			errs:  []error{ErrNodeTitleIsEmpty, ErrNodeSlugIsEmpty},
		},
		{
			desc:  "ForAnEmptySlug",
			slug:  "",
			title: "title",
			errs:  []error{ErrNodeSlugIsEmpty},
		},
		{
			desc:  "ForAnEmptyTitle",
			slug:  "slug",
			title: "",
			errs:  []error{ErrNodeTitleIsEmpty},
		},
		{
			desc:  "ForAnNonEmptyTitleAndSlug",
			slug:  "slug",
			title: "title",
			errs:  nil,
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			errs := Node{}.Validate(dto.Node{
				Content:    "",
				ContentURI: "",
				Kind:       "",
				Slug:       tc.slug,
				Title:      tc.title,
				Tags:       nil,
			})

			assert.Equal(t, tc.errs, errs)
		})
	}
}

func TestValidateNodeTitle(t *testing.T) {
	tcs := []struct {
		desc  string
		title string
		err   error
	}{
		{
			desc:  "ForAnEmptyTitle",
			title: "",
			err:   ErrNodeTitleIsEmpty,
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

			err := Node{}.validateTitle(tc.title)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestValidateNodeSlug(t *testing.T) {
	tcs := []struct {
		desc string
		slug string
		err  error
	}{
		{
			desc: "ForAnEmptyValue",
			slug: "",
			err:  ErrNodeSlugIsEmpty,
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

			err := Node{}.validateSlug(tc.slug)

			assert.Equal(t, tc.err, err)
		})
	}
}
