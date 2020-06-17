package config

// Database connection constants.
const (
	dbURI  = "bolt://localhost:7687/"
	dbUser = "neo4j"
	dbPwd  = "123456"
)

type Config struct {
	DB DB
}

func NewConfig() Config {
	return Config{
		DB: DB{
			URI:  dbURI,
			User: dbUser,
			Pwd:  dbPwd,
		},
	}
}