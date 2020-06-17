package dto

type Node struct {
	Content    string
	ContentURI string
	Kind       string
	Slug       string
	Tags       []string
	Title      string
}

// NewNode builds node DTO object from the functional options.
func NewNode(valueSetters ...NodeValueSetter) (node Node) {
	for _, vs := range valueSetters {
		vs(&node)
	}

	return node
}

// Node function option builders.
type NodeValueSetter func(o *Node)

func WithNodeContent(c string) NodeValueSetter    { return func(o *Node) { o.Content = c } }
func WithNodeContentURI(u string) NodeValueSetter { return func(o *Node) { o.ContentURI = u } }
func WithNodeKind(k string) NodeValueSetter       { return func(o *Node) { o.Kind = k } }
func WithNodeSlug(s string) NodeValueSetter       { return func(o *Node) { o.Slug = s } }
func WithNodeTags(tt []string) NodeValueSetter    { return func(o *Node) { o.Tags = tt } }
func WithNodeTitle(t string) NodeValueSetter      { return func(o *Node) { o.Title = t } }
