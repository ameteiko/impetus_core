package value

type ID struct {
	id string
}

func (id ID) String() string {
	return id.id
}