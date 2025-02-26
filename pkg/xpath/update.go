package xpath

import (
	"github.com/antchfx/xmlquery"
)

// UpdateNodeOrAttrByXPath updates the value of the node that matches the XPath expression in the XML document.
func UpdateNodeOrAttrByXPath(doc *xmlquery.Node, expr string, newValue string) bool {
	node := xmlquery.FindOne(doc, NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// Check if the expression is targeting an attribute
	if attrName, ok := GetAttributeNameFromExpression(expr); ok {
		if node.Parent == nil {
			return false
		}
		if updateAttribute(node.Parent, attrName, newValue) {
			return true
		}
		// Should we add new attributes?
		// xmlquery.AddAttr(node.Parent, attrName, newValue)
		return false
	}
	if node.FirstChild != nil {
		node.FirstChild.Data = newValue
	} else {
		node.Data = newValue
	}
	return true
}

// updateAttribute updates the value of the attribute with the given name in the node.
func updateAttribute(node *xmlquery.Node, attrName string, newValue string) bool {
	if !HasAttributes(node) {
		return false
	}
	for i, attr := range node.Attr {
		if attr.Name.Local == attrName {
			node.Attr[i].Value = newValue
			return true
		}
	}
	return false
}
