package tests

import (
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
)

func AssertContains(t *testing.T, doc *xmlquery.Node, expected string, showXML bool) {
	xml := doc.OutputXML(true) // stringify XML
	assert.NotEmpty(t, xml)
	assert.Contains(t, xml, expected)

	if showXML {
		t.Logf("XML: %s\n", xml)
	}
}

func AssertNotContains(t *testing.T, doc *xmlquery.Node, expected string, showXML bool) {
	xml := doc.OutputXML(true) // stringify XML
	assert.NotEmpty(t, xml)
	assert.NotContains(t, xml, expected)

	if showXML {
		t.Logf("XML: %s\n", xml)
	}
}
