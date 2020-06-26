package value

type (
	// TODO: Conceal all the interfaces for injected objects.
	Sanitiser interface {
		SanitiseID(string) string
		SanitiseSlug(string, string) string
		SanitiseTitle(string) string
	}

	Validator interface {
		ValidateID(string) error
		ValidateSlug(string) error
		ValidateTitle(string) error
	}

	Factory struct {
		sanitiser Sanitiser
		validator Validator
	}
)

func NewFactory(s Sanitiser, v Validator) Factory {
	return Factory{
		sanitiser: s,
		validator: v,
	}
}

func (f Factory) ID(id string) (ID, error) {
	id = f.sanitiser.SanitiseID(id)
	if err := f.validator.ValidateID(id); err != nil {
		return ID{}, err
	}

	return ID{id: id}, nil
}

// Slug constructs a slug Value Object from the slug value, or expectedTitle value if slug is empty.
func (f Factory) Slug(slug, title string) (Slug, error) {
	slug = f.sanitiser.SanitiseSlug(slug, title)
	if err := f.validator.ValidateSlug(slug); err != nil {
		return Slug{}, err
	}

	return Slug{slug: slug}, nil
}

func (f Factory) Title(t string) (Title, error) {
	t = f.sanitiser.SanitiseTitle(t)
	if err := f.validator.ValidateTitle(t); err != nil {
		return Title{}, err
	}

	return Title{title: t}, nil
}
