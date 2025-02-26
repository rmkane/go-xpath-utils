package xpathutils

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"
)

// AttributePathStart defines the prefix for attribute XPaths
const AttributePathStart = "/@"

// ParseXmlStr parses an XML string.
func ParseXmlStr(xmlStr string) (*xmlquery.Node, error) {
	return xmlquery.Parse(strings.NewReader(xmlStr))
}

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

// Serialize serializes the XML document into a string.
func Serialize(doc *xmlquery.Node) (string, error) {
	return doc.OutputXML(true), nil
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
	index, ok := getLastIndexOfLastAttribute(expr)
	if !ok {
		return "", false
	}
	return expr[index:], true
}

// getLastIndexOfLastAttribute extracts the attribute name from the expression.
func getLastIndexOfLastAttribute(expr string) (int, bool) {
	index := strings.LastIndex(expr, AttributePathStart)
	if index == -1 {
		return -1, false
	}
	return index + len(AttributePathStart), true
}

// RemoveAttributeFromXPath trims the trailing attribute selector from an XPath expression, if present.
// Example: "/root/item/@id" → "/root/item"
func RemoveAttributeFromXPath(expr string) (string, bool) {
	if !IsAttributeExpression(expr) {
		return expr, false
	}
	index := strings.LastIndex(expr, AttributePathStart)
	if index == -1 {
		return expr, false
	}
	return expr[:index], true
}
