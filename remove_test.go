package xpath

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmkane/go-xpath-utils/internal/tests"
	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

func TestRemoveNodeOrAttrByXPath_Node(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='25689f7cb-d9d3-47be-a734-667f0da8151b']"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := removeNodeOrAttrByXPath(doc, expr)
	assert.True(t, ok)

	// Verify that the node was removed
	tests.AssertNotContains(t, doc, "<blog id=\"25689f7cb-d9d3-47be-a734-667f0da8151b\">", false)
}

func TestRemoveNodeOrAttrByXPath_Attribute(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/@id"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := removeNodeOrAttrByXPath(doc, expr)
	assert.True(t, ok)

	// Verify that the attribute was removed
	tests.AssertNotContains(t, doc, "id=\"09c5e699-4d50-4da3-bd92-f37305c33ed4\"", false)
}
