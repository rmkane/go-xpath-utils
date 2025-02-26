package xpath

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmkane/go-xpath-utils/internal/tests"
	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

func TestAddNodeOrAttrByXPath_NodeValue(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs"
	key := "blog-foo"
	value := ""

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := addNodeOrAttrByXPath(doc, expr, key, value)
	assert.True(t, ok)

	tests.AssertContains(t, doc, "<blog-foo>", false)
}

func TestAddNodeOrAttrByXPath_AttributeValue(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/@updated"
	key := "updated"
	value := "42"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := addNodeOrAttrByXPath(doc, expr, key, value)
	assert.True(t, ok)

	tests.AssertContains(t, doc, "<blog id=\"09c5e699-4d50-4da3-bd92-f37305c33ed4\" updated=\"42\">", false)
}
