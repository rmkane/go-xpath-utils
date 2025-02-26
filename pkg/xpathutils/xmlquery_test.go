package xpathutils

import (
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
)

func TestChildren(t *testing.T) {
	t.Run("Has 3 children", func(t *testing.T) {
		node := &xmlquery.Node{Type: xmlquery.ElementNode}

		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.CommentNode})
		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.ElementNode})
		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.TextNode})

		children := Children(node)

		assert.Equal(t, 3, len(children))
	})
}

func TestChildNodes(t *testing.T) {
	t.Run("Has 1 child node", func(t *testing.T) {
		node := &xmlquery.Node{Type: xmlquery.ElementNode}

		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.CommentNode})
		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.ElementNode})
		xmlquery.AddChild(node, &xmlquery.Node{Type: xmlquery.TextNode})

		children, err := ChildNodes(node)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(children))
	})

	t.Run("Cannot have child nodes", func(t *testing.T) {
		node := &xmlquery.Node{Type: xmlquery.CommentNode}

		children, err := ChildNodes(node)

		assert.Error(t, err)
		assert.True(t, len(children) == 0)
	})
}

func TestIsElementLike(t *testing.T) {
	for _, test := range []struct {
		name     string
		n        *xmlquery.Node
		expected bool
	}{
		{
			name:     "DocumentNode can have children",
			n:        &xmlquery.Node{Type: xmlquery.DocumentNode},
			expected: true,
		},
		{
			name:     "DeclarationNode can have children",
			n:        &xmlquery.Node{Type: xmlquery.DeclarationNode},
			expected: true,
		},
		{
			name:     "ElementNode can have children",
			n:        &xmlquery.Node{Type: xmlquery.ElementNode},
			expected: true,
		},
		{
			name:     "TextNode cannot have children",
			n:        &xmlquery.Node{Type: xmlquery.TextNode},
			expected: false,
		},
		{
			name:     "CharDataNode cannot have children",
			n:        &xmlquery.Node{Type: xmlquery.CharDataNode},
			expected: false,
		},
		{
			name:     "CommentNode cannot have children",
			n:        &xmlquery.Node{Type: xmlquery.CommentNode},
			expected: false,
		},
		{
			name:     "AttributeNode cannot have children",
			n:        &xmlquery.Node{Type: xmlquery.AttributeNode},
			expected: false,
		},
		{
			name:     "NotationNode cannot have children",
			n:        &xmlquery.Node{Type: xmlquery.NotationNode},
			expected: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsElementLike(test.n))
		})
	}
}
