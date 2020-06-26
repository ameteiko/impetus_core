package value

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate mockery --name=sanitiser --output=. --inpackage --filename=factory_mock_sanitiser.go
//go:generate mockery --name=validator --output=. --inpackage --filename=factory_mock_validator.go

func TestID(t *testing.T) {
	tcs := []struct {
		testcase        string
		value           string
		sanitisedValue  string
		validationError error
		expectedID      ID
	}{
		{
			testcase:        "ForAnEmptyValue",
			value:           "",
			sanitisedValue:  "",
			validationError: nil,
			expectedID:      ID{},
		},
		{
			testcase:        "ForAValidValue",
			value:           "expectedID-value",
			sanitisedValue:  "expectedID-value",
			validationError: nil,
			expectedID:      ID{id: "expectedID-value"},
		},
		{
			testcase:        "ForAnInvalidValue",
			value:           "!!!",
			sanitisedValue:  "!!!",
			validationError: errors.New("invalid expectedID value"),
			expectedID:      ID{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.testcase, func(t *testing.T) {
			sanitiserMock := new(mockSanitiser)
			sanitiserMock.On("SanitiseID", tc.value).Return(tc.sanitisedValue).Once()
			validatorMock := new(mockValidator)
			validatorMock.On("ValidateID", tc.sanitisedValue).Return(tc.validationError).Once()

			id, err := NewFactory(sanitiserMock, validatorMock).ID(tc.value)

			assert.Equal(t, tc.expectedID, id)
			assert.Equal(t, err, tc.validationError)
			sanitiserMock.AssertExpectations(t)
			validatorMock.AssertExpectations(t)
		})
	}
}

func TestSlug(t *testing.T) {
	tcs := []struct {
		testcase        string
		value           string
		defaultValue    string
		sanitisedValue  string
		validationError error
		expectedSlug    Slug
	}{
		{
			testcase:        "ForAnEmptyValue",
			value:           "",
			defaultValue:    "",
			sanitisedValue:  "",
			validationError: nil,
			expectedSlug:    Slug{},
		},
		{
			testcase:        "ForAValidValue",
			value:           "Title",
			defaultValue:    "",
			sanitisedValue:  "Title",
			validationError: nil,
			expectedSlug:    Slug{slug: "Title"},
		},
		{
			testcase:        "ForAnInvalidValue",
			value:           "   ",
			defaultValue:    "",
			sanitisedValue:  "",
			validationError: errors.New("expectedTitle is empty"),
			expectedSlug:    Slug{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.testcase, func(t *testing.T) {
			sanitiserMock := new(mockSanitiser)
			sanitiserMock.On("SanitiseSlug", tc.value, tc.defaultValue).Return(tc.sanitisedValue).Once()
			validatorMock := new(mockValidator)
			validatorMock.On("ValidateSlug", tc.sanitisedValue).Return(tc.validationError).Once()

			slug, err := NewFactory(sanitiserMock, validatorMock).Slug(tc.value, tc.defaultValue)

			assert.Equal(t, tc.expectedSlug, slug)
			assert.Equal(t, err, tc.validationError)
			sanitiserMock.AssertExpectations(t)
			validatorMock.AssertExpectations(t)
		})
	}
}

func TestTitle(t *testing.T) {
	tcs := []struct {
		testcase        string
		value           string
		sanitisedValue  string
		validationError error
		expectedTitle   Title
	}{
		{
			testcase:        "ForAnEmptyValue",
			value:           "",
			sanitisedValue:  "",
			validationError: nil,
			expectedTitle:   Title{},
		},
		{
			testcase:        "ForAValidValue",
			value:           "Title",
			sanitisedValue:  "Title",
			validationError: nil,
			expectedTitle:   Title{title: "Title"},
		},
		{
			testcase:        "ForAnInvalidValue",
			value:           "   ",
			sanitisedValue:  "",
			validationError: errors.New("expectedTitle is empty"),
			expectedTitle:   Title{},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.testcase, func(t *testing.T) {
			sanitiserMock := new(mockSanitiser)
			sanitiserMock.On("SanitiseTitle", tc.value).Return(tc.sanitisedValue).Once()
			validatorMock := new(mockValidator)
			validatorMock.On("ValidateTitle", tc.sanitisedValue).Return(tc.validationError).Once()

			title, err := NewFactory(sanitiserMock, validatorMock).Title(tc.value)

			assert.Equal(t, tc.expectedTitle, title)
			assert.Equal(t, err, tc.validationError)
			sanitiserMock.AssertExpectations(t)
			validatorMock.AssertExpectations(t)
		})
	}
}
