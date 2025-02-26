package xpath

import (
	"fmt"
	"github.com/antchfx/xmlquery"

	"github.com/rmkane/go-xpath-utils/pkg/xpathutils"
)

// UpdateByXPathFromFile updates the value of a node or attribute that matches the XPath expression in the XML document.
func UpdateByXPathFromFile(filename, expr, newValue string) error {
	doc, err := xpathutils.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := updateNodeOrAttrByXPath(doc, expr, newValue); !ok {
		return fmt.Errorf("could not update node or attribute")
	}

	return xpathutils.SaveXML(doc, filename)
}

// UpdateByXPathFromString updates the value of a node or attribute that matches the XPath expression in the XML document.
func UpdateByXPathFromString(xmlStr, expr, newValue string) (string, error) {
	doc, err := xpathutils.ParseXmlStr(xmlStr)
	if err != nil {
		return "", err
	}

	if ok := updateNodeOrAttrByXPath(doc, expr, newValue); !ok {
		return "", fmt.Errorf("could not update node or attribute")
	}

	return xpathutils.Serialize(doc)
}

// updateNodeOrAttrByXPath updates the value of a node or attribute that matches the XPath expression in the given XML node tree.
func updateNodeOrAttrByXPath(doc *xmlquery.Node, expr string, newValue string) bool {
	node := xmlquery.FindOne(doc, xpathutils.NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// Check if the expression is targeting an attribute
	if attrName, ok := xpathutils.GetAttributeNameFromExpression(expr); ok {
		return xpathutils.SetAttrSafe(node.Parent, attrName, newValue)
	}

	// If the XPath is targeting an element, replace its data
	if node.FirstChild != nil {
		node.FirstChild.Data = newValue
	} else {
		node.Data = newValue
	}
	return true
}
