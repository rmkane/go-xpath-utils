package xpath

import (
	"fmt"

	"github.com/rmkane/go-xpath-utils/pkg/xpath"
)

// DropByXPath removes the node that matches the XPath expression from the XML document.
func DropByXPath(filename string, expr string) error {
	doc, err := xpath.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := xpath.DropNodeOrAttrByXPath(doc, expr); !ok {
		return fmt.Errorf("could not drop node")
	}

	return xpath.SaveXML(doc, filename)
}

// UpdateByXPath updates the value of the node that matches the XPath expression in the XML document.
func UpdateByXPath(filename string, expr string, newValue string) error {
	doc, err := xpath.LoadXML(filename)
	if err != nil {
		return err
	}

	if ok := xpath.UpdateNodeOrAttrByXPath(doc, expr, newValue); !ok {
		return fmt.Errorf("could not update value")
	}

	return xpath.SaveXML(doc, filename)
}
