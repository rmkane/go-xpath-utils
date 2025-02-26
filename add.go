package xpath

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// AddByXPathFromFile adds a node or attribute at the XPath location in the XML document.
// If the XPath targets an attribute, it will be added to the parent node.
// If the XPath targets an element, a new child node with the specified value will be created.
func AddByXPathFromFile(inputFile, outputFile, expr, key, value string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := addNodeOrAttrByXPath(doc, expr, key, value); !ok {
		return fmt.Errorf("failed to add node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
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
	trimmedExpr, _ := xpathutils.RemoveAttributeFromXPath(expr)

	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(trimmedExpr))
	if node == nil {
		return false
	}

	// If XPath targets an attribute, add it to the parent node
	if attrName != "" {
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
	return &xmlquery.Node{
		Type: xmlquery.ElementNode,
		Data: key,
		FirstChild: &xmlquery.Node{
			Type: xmlquery.TextNode,
			Data: value,
		},
	}
}
