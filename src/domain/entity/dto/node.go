package dto

type Node struct {
	ID string
	Content    string
	ContentURI string
	Kind       string
	Slug       string
	Tags       []string
	Title      string
}

// NewNode builds node DTO object from the functional options.
func NewNode(valueSetters ...nodeValueSetter) (node Node) {
	for _, vs := range valueSetters {
		vs(&node)
	}

	return node
}

func (n Node) IsZero() bool {
	zeroN := Node{}
	return n.Content == zeroN.Content &&
		n.ContentURI == zeroN.ContentURI &&
		n.Kind == zeroN.Kind &&
		n.Slug == zeroN.Slug &&
		n.Title == zeroN.Title &&
		(n.Tags == nil || len(n.Tags) == 0)
}

// Node function option builders.
type nodeValueSetter func(o *Node)

func WithNodeContent(c string) nodeValueSetter    { return func(o *Node) { o.Content = c } }
func WithNodeContentURI(u string) nodeValueSetter { return func(o *Node) { o.ContentURI = u } }
func WithNodeKind(k string) nodeValueSetter       { return func(o *Node) { o.Kind = k } }
func WithNodeSlug(s string) nodeValueSetter       { return func(o *Node) { o.Slug = s } }
func WithNodeTags(tt []string) nodeValueSetter    { return func(o *Node) { o.Tags = tt } }
func WithNodeTitle(t string) nodeValueSetter      { return func(o *Node) { o.Title = t } }
