package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	tcs := []struct {
		desc    string
		setters []nodeValueSetter
		node    Node
	}{
		{
			desc:    "ForAnEmptyConstructor",
			setters: nil,
			node:    Node{},
		},
		{
			desc:    "ForASetterOverride",
			setters: []nodeValueSetter{
				WithNodeTitle("First Title"),
				WithNodeTitle("Overwritten Title"),
			},
			node:    Node{
				Title: "Overwritten Title",
			},
		},
		{
			desc: "ForEverySetter",
			setters: []nodeValueSetter{
				WithNodeContent("content"),
				WithNodeContentURI("content/uri"),
				WithNodeKind("kind"),
				WithNodeSlug("slug"),
				WithNodeTitle("Title"),
				WithNodeTags([]string{"tags", "for", "node"}),
			},
			node: Node{
				Content:    "content",
				ContentURI: "content/uri",
				Kind:       "kind",
				Slug:       "slug",
				Tags:       []string{"tags", "for", "node"},
				Title:      "Title",
			},
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {

			var n Node
			if tc.setters == nil {
				n = NewNode()
			} else {
				n = NewNode(tc.setters...)
			}

			assert.Equal(t, tc.node, n)
		})
	}
}
