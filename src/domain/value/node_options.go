package value

type (
	nodeOptioner func(o *nodeOptions)
	nodeOptions  struct {
		kind string
		name string
	}
)

