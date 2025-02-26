package xpath

import (
	"github.com/antchfx/xmlquery"
)

// DropNodeOrAttrByXPath removes the node that matches the XPath expression from the XML document.
func DropNodeOrAttrByXPath(doc *xmlquery.Node, expr string) bool {
	node := xmlquery.FindOne(doc, NormalizeXPath(expr))
	if node == nil {
		return false
	}

	// Check if the expression is targeting an attribute
	if attrName, ok := GetAttributeNameFromExpression(expr); ok {
		if node.Parent == nil {
			return false
		}
		if removeAttribute(node.Parent, attrName) {
			return true
		}
		return false
	}

	return RemoveXMLNode(node)
}

func removeAttribute(node *xmlquery.Node, attrName string) bool {
	if !HasAttributes(node) {
		return false
	}
	for i, attr := range node.Attr {
		if attr.Name.Local == attrName {
			node.Attr = append(node.Attr[:i], node.Attr[i+1:]...)
			return true
		}
	}
	return false
}
