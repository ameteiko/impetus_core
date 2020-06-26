package dto

type Relation struct {
	ID   string
	Kind string
}

// NewNode builds node DTO object from the functional options.
func NewRelation(valueSetters ...relationValueSetter) (relation Relation) {
	for _, vs := range valueSetters {
		vs(&relation)
	}

	return relation
}

func (r Relation) IsZero() bool {
	return r == Relation{}
}

// Node function option builders.
type relationValueSetter func(r *Relation)

func WithRelationID(id string) relationValueSetter  { return func(r *Relation) { r.ID = id } }
func WithRelationKind(k string) relationValueSetter { return func(r *Relation) { r.Kind = k } }
