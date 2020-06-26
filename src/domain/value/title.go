package value

type Title struct {
	title string
}

func (t Title) String() string {
	return t.title
}
