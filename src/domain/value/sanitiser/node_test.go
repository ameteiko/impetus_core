package sanitiser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ameteiko/mindnet/src/domain/value/dto"
)

func TestNodeSanitise(t *testing.T) {
	tcs := []struct {
		desc          string
		rawNode       dto.Node
		sanitisedNode dto.Node
	}{
		{
			desc:          "TestSanitiseForAnEmptyDTO",
			rawNode:       dto.Node{},
			sanitisedNode: dto.Node{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {
			s := Node{}

			sanitisedNode := s.Sanitise(tc.rawNode)

			assert.Equal(t, tc.sanitisedNode, sanitisedNode)
		})
	}
}

func TestSanitiseNodeTitle(t *testing.T) {
	tcs := []struct {
		desc           string
		title          string
		sanitisedTitle string
	}{
		{
			desc:           "ForAnEmptyValue",
			title:          "",
			sanitisedTitle: "",
		},
		{
			desc:           "ForAWhitespacesValue",
			title:          "   ",
			sanitisedTitle: "",
		},
		{
			desc:           "ForALowerCaseValue",
			title:          "scripting",
			sanitisedTitle: "scripting",
		},
		{
			desc:           "ForAWhitespacesBetweenWords",
			title:          " sheLL   sCRipting",
			sanitisedTitle: "sheLL sCRipting",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			sanitisedTitle := Node{}.sanitiseTitle(tc.title)

			assert.Equal(t, tc.sanitisedTitle, sanitisedTitle)
		})
	}
}

func TestSanitiseNodeKind(t *testing.T) {
	tcs := []struct {
		desc          string
		kind          string
		sanitisedKind string
	}{
		{
			desc:          "ForAnEmptyValue",
			kind:          "",
			sanitisedKind: "",
		},
		{
			desc:          "ForAnEmptyWhitespacesKind",
			kind:          "   ",
			sanitisedKind: "",
		},
		{
			desc:          "ForAValidKindValue",
			kind:          " artefact  ",
			sanitisedKind: "artefact",
		},
		{
			desc:          "ForARegisteredKindValueShortcut",
			kind:          "artf  ",
			sanitisedKind: "artefact",
		},
		{
			desc:          "ForARegisteredMisspelledShortcut",
			kind:          "artifact",
			sanitisedKind: "artefact",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			sanitisedKind := Node{}.sanitiseKind(tc.kind)

			assert.Equal(t, tc.sanitisedKind, sanitisedKind)
		})
	}
}

func TestSanitiseNodeSlug(t *testing.T) {
	tcs := []struct {
		desc          string
		slug          string
		title         string
		sanitisedSlug string
	}{
		{
			desc:          "ForACleanSlug",
			slug:          "shell",
			title:         "",
			sanitisedSlug: "shell",
		},
		{
			desc:          "ForACleanSlugAndNotEmptyTitle",
			slug:          "shell",
			title:         "title",
			sanitisedSlug: "shell",
		},
		{
			desc:          "ForAnEmptySlug",
			slug:          "",
			title:         "",
			sanitisedSlug: "",
		},
		{
			desc:          "ForAnEmptySlugAndNotEmptyTitle",
			slug:          "",
			title:         " title ",
			sanitisedSlug: "title",
		},
		{
			desc:          "ForAnEmptySlugAndNotEmptyMultiWordTitle",
			slug:          "",
			title:         "  SheLL  ScripTing env    variables    ",
			sanitisedSlug: "shell-scripting-env-variables",
		},
		{
			desc:          "ForAnEmptyWhitespacesSlug",
			slug:          "  ",
			title:         "",
			sanitisedSlug: "",
		},
		{
			desc:          "ForASlugWrappedWithWhitespaces",
			slug:          " shell  ",
			title:         "",
			sanitisedSlug: "shell",
		},
		{
			desc:          "ForASlugWithCapitalLetters",
			slug:          "SheLL",
			title:         "",
			sanitisedSlug: "shell",
		},

		{
			desc:          "ForASlugWithSeveralWords",
			slug:          " SheLL  ScripTing env    variables   ",
			title:         "",
			sanitisedSlug: "shell-scripting-env-variables",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			sanitisedSlug := Node{}.sanitiseSlug(tc.slug, tc.title)

			assert.Equal(t, tc.sanitisedSlug, sanitisedSlug)
		})
	}
}

func TestSanitiseNodeTags(t *testing.T) {
	tcs := []struct {
		desc          string
		tags          []string
		sanitisedTags []string
	}{
		{
			desc:          "ForANilTagsList",
			tags:          nil,
			sanitisedTags: nil,
		},
		{
			desc:          "ForAnEmptyTagsList",
			tags:          []string{},
			sanitisedTags: nil,
		},
		{
			desc:          "ForASingleTag",
			tags:          []string{"  tag   "},
			sanitisedTags: []string{"tag"},
		},
		{
			desc:          "ForASingleMultiWordTagTag",
			tags:          []string{"  tag  value    "},
			sanitisedTags: []string{"tag value"},
		},
		{
			desc:          "ForSeveralIdenticalTags",
			tags:          []string{"  some tag", "some tag  "},
			sanitisedTags: []string{"some tag"},
		},
		{
			desc:          "ForMultipleTags",
			tags:          []string{"  some tag", "yet  another tag  ", "   some   tag  "},
			sanitisedTags: []string{"some tag", "yet another tag"},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			sanitisedTags := Node{}.sanitiseTags(tc.tags)

			assert.Equal(t, tc.sanitisedTags, sanitisedTags)
		})
	}
}
