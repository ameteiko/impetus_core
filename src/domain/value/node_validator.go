package value

// nodeKindMapping is a map for all acceptable values for each node kind.
var nodeKindMapping = map[string][]string{
	NodeArtefact.String(): {"artifact", "artf", "art", "atf"},
	NodeSnippet.String(): {"snipet", "snpt", "snp"},
}

type NodeValidator struct{}

func NewNodeValidator() NodeValidator {
	return NodeValidator{}
}

func ValidateKind(k string) error {
	if k == "" {
		return nil
	}



	// K is in registered values

	return nil
}
