package xpath

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
)

// LoadXML loads an XML file and returns the root node.
func LoadXML(filename string) (*xmlquery.Node, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(xmlFile *os.File) {
		err = xmlFile.Close()
		if err != nil {
			slog.Error(fmt.Sprintf("Error closing XML file: %v", err))
		}
	}(xmlFile)

	doc, err := xmlquery.Parse(xmlFile)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// SaveXML saves the XML document to a file.
func SaveXML(doc *xmlquery.Node, filename string) error {
	if doc == nil {
		return fmt.Errorf("document is nil")
	}
	if filename == "" {
		return fmt.Errorf("filename is empty")
	}
	xml := doc.OutputXML(true)
	return os.WriteFile(filename, []byte(xml), 0644)
}

// NormalizeXPath ensures that the XPath expression starts with a double slash.
func NormalizeXPath(expr string) string {
	return "//" + strings.TrimLeft(expr, "/")
}

// IsAttributeExpression checks if the expression is targeting an attribute.
func IsAttributeExpression(expr string) bool {
	return regexp.MustCompile(`/@[\w-]+$`).MatchString(expr)
}

// GetAttributeNameFromExpression extracts the attribute name from the expression.
func GetAttributeNameFromExpression(expr string) (string, bool) {
	if !IsAttributeExpression(expr) {
		return "", false
	}
	return expr[strings.LastIndex(expr, "/@")+2:], true
}

// HasAttributes checks if the node has any attributes.
func HasAttributes(node *xmlquery.Node) bool {
	return node != nil && len(node.Attr) > 0
}

// RemoveXMLNode removes the node from the XML document.
func RemoveXMLNode(node *xmlquery.Node) bool {
	if node == nil || node.Parent == nil {
		return false
	}

	parent := node.Parent

	if node.PrevSibling != nil {
		node.PrevSibling.NextSibling = node.NextSibling
	} else {
		parent.FirstChild = node.NextSibling
	}

	if node.NextSibling != nil {
		node.NextSibling.PrevSibling = node.PrevSibling
	} else {
		parent.LastChild = node.PrevSibling
	}

	node.Parent = nil
	node.PrevSibling = nil
	node.NextSibling = nil

	return true
}
