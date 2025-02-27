package xpathutils

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// GetAttrByXPathFromFile retrieves the value of an attribute at the specified XPath in the XML document.
func GetAttrByXPathFromFile(inputFile, expr string) (string, error) {
	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return "", err
	}

	value, ok := getAttrByXPath(doc, expr)
	if !ok {
		return "", fmt.Errorf("failed to get node or attribute at XPath: %s", expr)
	}
	return value, nil
}

// GetAttrByXPathFromString retrieves the value of an attribute at the specified XPath in the given XML string.
func GetAttrByXPathFromString(xmlStr, expr string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	value, ok := getAttrByXPath(doc, expr)
	if !ok {
		return "", fmt.Errorf("failed to get node or attribute at XPath: %s", expr)
	}
	return value, nil
}

// getAttrByXPath retrieves the value of an attribute at the specified XPath in the given XML node tree.
func getAttrByXPath(doc *xmlquery.Node, expr string) (string, bool) {
	attrName, _ := xpathutils.GetAttributeNameFromExpression(expr)
	trimmedExpr, _ := xpathutils.RemoveAttributeFromXPath(expr)

	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(trimmedExpr))
	if node == nil {
		return "", false
	}

	// If XPath targets an attribute, add it to the parent node
	if attrName == "" {
		return "", false
	}

	return xpathutils.GetAttrSafe(node, attrName)
}
