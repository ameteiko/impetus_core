package sanitiser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	tcs := []struct {
		desc         string
		value        string
		trimmedValue string
	}{
		{
			desc:         "ForAnEmptyValue",
			value:        "",
			trimmedValue: "",
		},
		{
			desc:         "ForAnEmptyWhitespacesValue",
			value:        "  ",
			trimmedValue: "",
		},
		{
			desc:         "ForAValueWrappedWithWhitespaces",
			value:        " shell  ",
			trimmedValue: "shell",
		},
		{
			desc:         "ForAMultiLineValueWrappedWithWhitespaces",
			value:        "  some\nmulti-lines content\n GOes  \nhere ",
			trimmedValue: "some\nmulti-lines content\n GOes  \nhere",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			trimmedValue := trim(tc.value)

			assert.Equal(t, tc.trimmedValue, trimmedValue)
		})
	}
}

func TestMergeWhitespaces(t *testing.T) {
	tcs := []struct {
		desc           string
		value          string
		sanitisedValue string
	}{
		{
			desc:           "ForAnEmptyValue",
			value:          "",
			sanitisedValue: "",
		},
		{
			desc:           "ForAWhitespacesValue",
			value:          "   ",
			sanitisedValue: "",
		},
		{
			desc:           "ForACleanValue",
			value:          "scripting",
			sanitisedValue: "scripting",
		},
		{
			desc:           "ForAWhitespacesBetweenWords",
			value:          " sheLL   sCRipting   env  v",
			sanitisedValue: "sheLL sCRipting env v",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			sanitisedValue := mergeWhitespaces(tc.value)

			assert.Equal(t, tc.sanitisedValue, sanitisedValue)
		})
	}
}
