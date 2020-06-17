package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
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

func (db *DB) CreateNode(label string, slug string, contentURI string) error {
	result, err := db.session.Run(fmt.Sprintf(`CREATE (n: %s { slug: $slug, content_uri: $contentURI}) RETURN n`, label),
		map[string]interface{}{
			"label":      label,
			"slug":       slug,
			"contentURI": contentURI,
		})

	if err != nil {
		return err
	}

	for result.Next() {
		fmt.Println(result)
		// fmt.Printf("Created Item with Id = '%d' and Slug = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	}
	if err = result.Err(); err != nil {
		return err // handle error
	}

	return err
}

func (db *DB) ConnectNodes(slugA, slugB string, relationType string) {

}
