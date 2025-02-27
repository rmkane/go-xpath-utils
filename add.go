package xpathutils

import (
	"fmt"

	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// AddAttrByXPathFromFile adds an attribute at the XPath location in the XML document.
func AddAttrByXPathFromFile(inputFile, outputFile, expr, value string) error {
	if outputFile == "" {
		outputFile = inputFile
	}

	doc, err := xpathutils.LoadXML(inputFile)
	if err != nil {
		return err
	}

	if ok := addAttrByXPath(doc, expr, value); !ok {
		return fmt.Errorf("failed to add node or attribute at XPath: %s", expr)
	}

	return xpathutils.SaveXML(doc, outputFile)
}

// AddAttrByXPathFromString adds an attribute at the XPath location in the given XML string.
func AddAttrByXPathFromString(xmlStr, expr, value string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := addAttrByXPath(doc, expr, value); !ok {
		return "", fmt.Errorf("failed to add node or attribute at XPath: %s", expr)
	}

	return xpathutils.Serialize(doc)
}

// addAttrByXPath adds an attribute at the specified XPath in the given XML node tree.
func addAttrByXPath(doc *xmlquery.Node, expr, value string) bool {
	attrName, _ := xpathutils.GetAttributeNameFromExpression(expr)
	if attrName == "" {
		return false
	}

	trimmedExpr, _ := xpathutils.RemoveAttributeFromXPath(expr)

	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(trimmedExpr))
	if node == nil {
		return false
	}

	return xpathutils.AddAttrSafe(node, attrName, value)
}
