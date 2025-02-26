package xpath

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// RemoveByXPathFromFile removes a node or attribute at the specified XPath from the XML document.
// If the XPath targets an attribute, it will be removed from the parent node.
// If the XPath targets an element, the node and its subtree will be removed.
func RemoveByXPathFromFile(filename, expr string) error {
	doc, err := xpathutils.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := removeNodeOrAttrByXPath(doc, expr); !ok {
		return fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, filename)
}

// RemoveByXPathFromString removes a node or attribute at the specified XPath from the given XML string.
// Returns the updated XML string or an error if the operation fails.
func RemoveByXPathFromString(xmlStr, expr string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := removeNodeOrAttrByXPath(doc, expr); !ok {
		return "", fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// removeNodeOrAttrByXPath removes a node or attribute at the specified XPath in the given XML node tree.
// - If the XPath targets an attribute, it is removed from the parent node.
// - If the XPath targets an element, the entire node (including children) is removed.
func removeNodeOrAttrByXPath(doc *xmlquery.Node, expr string) bool {
	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// If XPath targets an attribute, remove it from the parent node
	if attrName, ok := xpathutils.GetAttributeNameFromExpression(expr); ok {
		if node.Parent == nil {
			return false // Cannot remove an attribute without a parent node
		}
		return xpathutils.RemoveAttrSafe(node.Parent, attrName)
	}

	// If XPath targets an element, remove the node from the tree
	return xpathutils.RemoveFromTreeSafe(node)
}
