package xpathutils

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

//
// ─── NODE TYPE HELPERS ─────────────────────────────────────────────────────────
//

// IsElementLike returns true if the node type supports child nodes.
func IsElementLike(node *xmlquery.Node) bool {
	if node == nil {
		return false
	}
	switch node.Type {
	case xmlquery.ElementNode, xmlquery.DocumentNode, xmlquery.DeclarationNode:
		return true
	default:
		return false
	}
}

// HasAttributes checks if the node has any attributes.
func HasAttributes(node *xmlquery.Node) bool {
	return node != nil && len(node.Attr) > 0
}

// HasAttr determines if a node has an attribute with the given key.
func HasAttr(node *xmlquery.Node, key string) bool {
	if !HasAttributes(node) {
		return false
	}
	name := newXMLName(key)
	for _, attr := range node.Attr {
		if attr.Name == name {
			return true
		}
	}
	return false
}

//
// ─── CHILD NODE HELPERS ─────────────────────────────────────────────────────────
//

// Children returns a slice of all direct child nodes of the given node.
// This includes all node types that are children, without filtering.
func Children(node *xmlquery.Node) []*xmlquery.Node {
	if node == nil {
		return nil
	}
	var children []*xmlquery.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}
	return children
}

// ChildNodes returns a slice of all direct child nodes of the given node,
// excluding attributes, text, comments, and char data.
// Returns an error if the node type does not support child nodes.
func ChildNodes(node *xmlquery.Node) ([]*xmlquery.Node, error) {
	if !IsElementLike(node) {
		return nil, fmt.Errorf("node type %v does not support child nodes", node.Type)
	}
	var children []*xmlquery.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if IsElementLike(child) {
			children = append(children, child)
		}
	}
	return children, nil
}

//
// ─── NODE MODIFICATION HELPERS ──────────────────────────────────────────────────
//

// RemoveFromTreeSafe removes a node and its subtree from the document
// tree it is in. If the node is the root of the tree, then it's a no-op.
//
// Returns true if the node was removed, false if it was not.
func RemoveFromTreeSafe(node *xmlquery.Node) bool {
	if node == nil || node.Parent == nil {
		return false
	}
	xmlquery.RemoveFromTree(node)
	return true
}

// AddAttrSafe adds an attribute to the node only if it does not already exist.
//
// Returns true if the attribute was added, false if it already exists.
func AddAttrSafe(node *xmlquery.Node, key, value string) bool {
	if HasAttr(node, key) {
		return false
	}
	node.SetAttr(key, value)
	return true
}

// SetAttrSafe updates an existing attribute's value, or returns false if the attribute does not exist.
func SetAttrSafe(node *xmlquery.Node, key, value string) bool {
	if !HasAttr(node, key) { // Fixed check (was incorrectly checking 'value')
		return false
	}
	node.SetAttr(key, value)
	return true
}

// RemoveAttrSafe removes an attribute from a node if it exists.
//
// Returns true if the attribute was removed, false otherwise.
func RemoveAttrSafe(node *xmlquery.Node, key string) bool {
	if !HasAttr(node, key) {
		return false
	}
	node.RemoveAttr(key)
	return true
}

//
// ─── UTILITY FUNCTIONS ──────────────────────────────────────────────────────────
//

// newXMLName converts a string into a xml.Name struct.
// This is a copy of the private [xmlquery.newXMLName] function.
func newXMLName(name string) xml.Name {
	if i := strings.IndexByte(name, ':'); i > 0 {
		return xml.Name{
			Space: name[:i],
			Local: name[i+1:],
		}
	}
	return xml.Name{
		Local: name,
	}
}
