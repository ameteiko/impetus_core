package value

type Slug struct {
	slug string
}

func (s Slug) String() string {
	return s.slug
}