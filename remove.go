package xpathutils

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// RemoveAttrByXPathFromFile removes an attribute at the specified XPath from the XML document.
func RemoveAttrByXPathFromFile(inputFile, outputFile, expr string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := removeAttrByXPath(doc, expr); !ok {
		return fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
}

// RemoveAttrByXPathFromString removes an attribute at the specified XPath from the given XML string.
func RemoveAttrByXPathFromString(xmlStr, expr string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := removeAttrByXPath(doc, expr); !ok {
		return "", fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// RemoveNodeByXPathFromFile removes an attribute at the specified XPath from the XML document.
func RemoveNodeByXPathFromFile(inputFile, outputFile, expr string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := removeNodeByXPath(doc, expr); !ok {
		return fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
}

// RemoveNodeByXPathFromString removes an attribute at the specified XPath from the given XML string.
func RemoveNodeByXPathFromString(xmlStr, expr string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := removeNodeByXPath(doc, expr); !ok {
		return "", fmt.Errorf("failed to remove node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// removeAttrByXPath removes an attribute at the specified XPath in the given XML node tree.
func removeAttrByXPath(doc *xmlquery.Node, expr string) bool {
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

	return false
}

func removeNodeByXPath(doc *xmlquery.Node, expr string) bool {
	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	if xpathutils.IsAttributeExpression(expr) {
		return false
	}

	xmlquery.RemoveFromTree(node)
	return true
}
