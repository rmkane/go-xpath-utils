package xpathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmkane/go-xpath-utils/internal/tests"
	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

func TestUpdateAttrByXPath(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/author/@name"
	newValue := "Jim Doe"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := updateAttrByXPath(doc, expr, newValue)
	assert.True(t, ok)

	tests.AssertContains(t, doc, "<author name=\"Jim Doe\" posts=\"1\"></author>", false)
}
