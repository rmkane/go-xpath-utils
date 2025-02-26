package xpath

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmkane/go-xpath-utils/internal/tests"
	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

func TestUpdateNodeOrAttrByXPath_NodeValue(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='25689f7cb-d9d3-47be-a734-667f0da8151b']/posts/post[@id='302d48df-ceed-4b57-b0ff-12f17fe5759f']/title"
	newValue := "Second Post [Updated]"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := updateNodeOrAttrByXPath(doc, expr, newValue)
	assert.True(t, ok)

	tests.AssertContains(t, doc, "<title>Second Post [Updated]</title>", false)
}

func TestUpdateNodeOrAttrByXPath_AttributeValue(t *testing.T) {
	filename := "./testdata/obj.xml"
	expr := "/blogs/blog[@id='09c5e699-4d50-4da3-bd92-f37305c33ed4']/author/@name"
	newValue := "Jim Doe"

	doc, err := xpathutils.LoadXML(filename)
	assert.NoError(t, err)

	ok := updateNodeOrAttrByXPath(doc, expr, newValue)
	assert.True(t, ok)

	tests.AssertContains(t, doc, "<author name=\"Jim Doe\" posts=\"1\"></author>", false)
}
