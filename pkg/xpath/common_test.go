package xpath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeXPath(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"div[@id='test']", "//div[@id='test']"},
		{"/div[@id='test']", "//div[@id='test']"},
		{"///div[@id='test']", "//div[@id='test']"},
	}

	for _, tc := range testCases {
		actual := NormalizeXPath(tc.input)
		assert.Equal(t, tc.expected, actual, tc.input)
	}
}

func TestIsAttributeExpression(t *testing.T) {
	testCases := []struct {
		message    string
		expression string
		expected   bool
	}{
		{"valid attribute xpath", "//foo/@id", true},
		{"valid attribute xpath", "//foo/bar/@name", true},
		{"valid element xpath", "//div[@id='value']", false},
		{"invalid attribute xpath", "@foo", false},
	}

	for _, tc := range testCases {
		actual := IsAttributeExpression(tc.expression)
		assert.Equal(t, tc.expected, actual, tc.message)
	}
}

func TestGetAttributeNameFromExpression(t *testing.T) {
	testCases := []struct {
		expression string
		expected   string
		ok         bool
	}{
		{"//foo/@id", "id", true},
		{"//foo/bar/@name", "name", true},
		{"//div[@id='value']", "", false},
		{"@foo", "", false},
		{"//foo/@", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.expression, func(t *testing.T) {
			actual, ok := GetAttributeNameFromExpression(tc.expression)
			assert.Equal(t, tc.ok, ok)
			assert.Equal(t, tc.expected, actual, tc.expression)
		})
	}
}
