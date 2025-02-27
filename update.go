package xpath

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// UpdateByXPathFromFile updates the value of an attribute at the specified XPath in the XML document.
func UpdateByXPathFromFile(inputFile, outputFile, expr, newValue string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := updateAttrByXPath(doc, expr, newValue); !ok {
		return fmt.Errorf("failed to update node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
}

// UpdateByXPathFromString updates the value of an attribute at the specified XPath in the given XML string.
func UpdateByXPathFromString(xmlStr, expr, newValue string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := updateAttrByXPath(doc, expr, newValue); !ok {
		return "", fmt.Errorf("failed to update node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// updateNodeOrAttrByXPath updates the value of an attribute of a node at the specified XPath in the given XML node tree.
func updateAttrByXPath(doc *xmlquery.Node, expr, newValue string) bool {
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
	return false
}
