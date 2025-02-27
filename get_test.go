package xpathutils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

func TestGetAttrByXPath_NodeValue(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/@id"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	value, ok := getAttrByXPath(doc, expr)
	assert.True(t, ok)
	assert.True(t, value == "09c5e699-4d50-4da3-bd92-f37305c33ed4")
}
