package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/ameteiko/mindnet/domain/entity"
	"github.com/ameteiko/mindnet/domain/value"
)

const (
	fieldContent    = "content"
	fieldContentURI = "contentURI"
	fieldSlug       = "slug"
	fieldTitle      = "title"

	defaultNodeType = "artefact"
)

// DB is an secondary adapter for the database layer. Meaning that it represents a connection to the backend databases.
type DB struct {
	session neo4j.Session
}

func NewDB(dbSession neo4j.Session) *DB {
	return &DB{
		session: dbSession,
	}
}

func (db *DB) CreateNode(n entity.Node) error {
	fmt.Println(n.Slug, n.Title, n.ID)
	result, err := db.session.Run(
		fmt.Sprintf(`CREATE (n: %s { slug: $slug, title: $title}) RETURN n`, defaultNodeType),
		map[string]interface{}{
			fieldSlug:  n.Slug.String(),
			fieldTitle: n.Title.String(),
		},
	)
	if err != nil {
		return err
	}

	for result.Next() {
		// Log it here.
		// fmt.Println(result)
		// fmt.Printf("Created Item with Id = '%d' and Slug = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	}

	return result.Err()
}

func (db *DB) GetNodeBySlug(slug value.Slug) (node entity.Node, err error) {
	return node, err
}
