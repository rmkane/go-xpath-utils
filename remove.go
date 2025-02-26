package xpath

import (
	"fmt"
	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// RemoveByXPathFromFile removes the node or attribute that matches the XPath expression from the XML document.
func RemoveByXPathFromFile(filename, expr string) error {
	doc, err := xpathutils.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := removeNodeOrAttrByXPath(doc, expr); !ok {
		return fmt.Errorf("could not remove node or attribute")
	}

	return xpathutils.SaveXML(doc, filename)
}

// RemoveByXPathFromString removes the node or attribute that matches the XPath expression from the XML document.
func RemoveByXPathFromString(xmlStr, expr string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := removeNodeOrAttrByXPath(doc, expr); !ok {
		return "", fmt.Errorf("could not remove node or attribute")
	}

	return xpathutils.Serialize(doc)
}

// removeNodeOrAttrByXPath removes a node or attribute that matches the XPath expression in the given XML node tree.
func removeNodeOrAttrByXPath(doc *xmlquery.Node, expr string) bool {
	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// Check if the expression targets an attribute
	if attrName, ok := xpathutils.GetAttributeNameFromExpression(expr); ok {
		return xpathutils.RemoveAttrSafe(node.Parent, attrName)
	}

	return xpathutils.RemoveFromTreeSafe(node)
}
