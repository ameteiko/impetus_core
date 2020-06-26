package di

import (
	"fmt"
	
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4JSession returns a Neo4J database session object.
func Neo4JSession(dbURI, dbUser, dbPwd string) (neo4j.Session, error) {
	dbDriver, err := neo4j.NewDriver(dbURI, neo4j.BasicAuth(dbUser, dbPwd, ""))
	if err != nil {
		return nil, fmt.Errorf("cannot establish database connection: %w", err)
	}

	neo4JSession, err := dbDriver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, fmt.Errorf("cannot establish database connection: %w", err)
	}

	return neo4JSession, nil
}
