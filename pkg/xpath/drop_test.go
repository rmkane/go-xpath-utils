package xpath

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"xpath-example/internal/tests"
)

func TestDropNodeOrAttrByXPath_Node(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='25689f7cb-d9d3-47be-a734-667f0da8151b']"

	doc, err := LoadXML(filename)
	assert.NoError(t, err)

	ok := DropNodeOrAttrByXPath(doc, expr)
	assert.True(t, ok)

	// Verify that the node was removed
	tests.AssertNotContains(t, doc, "<blog id=\"25689f7cb-d9d3-47be-a734-667f0da8151b\">", false)
}

func TestDropNodeOrAttrByXPath_Attribute(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/@id"

	doc, err := LoadXML(filename)
	assert.NoError(t, err)

	ok := DropNodeOrAttrByXPath(doc, expr)
	assert.True(t, ok)

	// Verify that the attribute was removed
	tests.AssertNotContains(t, doc, "id=\"09c5e699-4d50-4da3-bd92-f37305c33ed4\"", false)
}
