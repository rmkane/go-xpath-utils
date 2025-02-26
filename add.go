package xpath

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// AddByXPathFromFile adds a node or attribute at the XPath location in the XML document.
// If the XPath targets an attribute, it will be added to the parent node.
// If the XPath targets an element, a new child node with the specified value will be created.
func AddByXPathFromFile(filename, expr, key, value string) error {
	doc, err := xpathutils.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := addNodeOrAttrByXPath(doc, expr, key, value); !ok {
		return fmt.Errorf("failed to add node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, filename)
}

// AddByXPathFromString adds a node or attribute at the XPath location in the given XML string.
// Returns the updated XML string or an error if the operation fails.
func AddByXPathFromString(xmlStr, expr, key, value string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := addNodeOrAttrByXPath(doc, expr, key, value); !ok {
		return "", fmt.Errorf("failed to add node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// addNodeOrAttrByXPath adds a node or attribute at the specified XPath location in the XML node tree.
// - If the XPath targets an attribute, it is added to the parent node.
// - If the XPath targets an element, a new child node is created with the given value.
func addNodeOrAttrByXPath(doc *xmlquery.Node, expr, key, value string) bool {
	attrName, _ := xpathutils.GetAttributeNameFromExpression(expr)
	isAttribute := attrName != ""

	if isAttribute {
		// Make sure we can find the element, so remove the attribute from the path
		expr = strings.TrimSuffix(expr, xpathutils.AttributePathStart+attrName)
	}

	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// If XPath targets an attribute, add it to the parent node
	if isAttribute {
		if attrName != key {
			return false
		}

		return xpathutils.AddAttrSafe(node, key, value)
	}

	// If XPath targets an element, add a new child node
	xmlquery.AddChild(node, createWrappedTextNode(key, value))
	return true
}

func createWrappedTextNode(key, value string) *xmlquery.Node {
	node := xmlquery.Node{
		Type: xmlquery.ElementNode,
		Data: key,
	}
	xmlquery.AddChild(&node, &xmlquery.Node{
		Type: xmlquery.TextNode,
		Data: value,
	})
	return &node
}
