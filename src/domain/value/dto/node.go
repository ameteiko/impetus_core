package dto

type (
	Node struct {
		Title string
		Slug  string
		Kind  string
	}

	NodeChanger func(o *Node)
)

// NewNode builds node DTO object from the functional options.
func NewNode(changers ...NodeChanger) (node Node) {
	for _, c := range changers {
		c(&node)
	}

	return node
}

// WithNodeTitle is a node creation functional option.
func WithNodeTitle(t string) NodeChanger {
	return func(o *Node) {
		o.Title = t
	}
}
// WithNodeName is a node creation functional option.
func WithNodeSlug(slug string) NodeChanger {
	return func(o *Node) {
		o.Slug = slug
	}
}

// WithNodeKind is a node creation functional option.
func WithNodeKind(k string) NodeChanger {
	return func(o *Node) {
		o.Kind = k
	}
}
