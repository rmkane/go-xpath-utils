package xpath

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// UpdateByXPathFromFile updates the value of a node or attribute at the specified XPath in the XML document.
// If the XPath targets an attribute, its value is updated.
// If the XPath targets an element, its content is replaced.
func UpdateByXPathFromFile(inputFile, outputFile, expr, newValue string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := updateNodeOrAttrByXPath(doc, expr, newValue); !ok {
		return fmt.Errorf("failed to update node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
}

// UpdateByXPathFromString updates the value of a node or attribute at the specified XPath in the XML string.
// Returns the updated XML string or an error if the operation fails.
func UpdateByXPathFromString(xmlStr, expr, newValue string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := updateNodeOrAttrByXPath(doc, expr, newValue); !ok {
		return "", fmt.Errorf("failed to update node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// updateNodeOrAttrByXPath updates the value of a node or attribute at the specified XPath in the given XML node tree.
// - If the XPath targets an attribute, it is updated in the parent node.
// - If the XPath targets an element, its content is replaced.
func updateNodeOrAttrByXPath(doc *xmlquery.Node, expr, newValue string) bool {
	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// If XPath targets an attribute, update it in the parent node
	if attrName, ok := xpathutils.GetAttributeNameFromExpression(expr); ok {
		if node.Parent == nil {
			return false // Cannot update an attribute without a parent node
		}
		return xpathutils.SetAttrSafe(node.Parent, attrName, newValue)
	}

	// If XPath targets an element, replace its content
	if node.FirstChild != nil {
		node.FirstChild.Data = newValue
	} else {
		node.Data = newValue
	}
	return true
}
