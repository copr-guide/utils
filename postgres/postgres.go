package postgres

import (
	"database/sql"

	database "github.com/copr-guide/utils/internal/database"
	logger "github.com/copr-guide/utils/log"
)

type Pg struct {
	DB *database.Queries
}

func CreatePostgresConn(dbUrl string) *Pg {
	if dbUrl == "" {
		logger.LogFatal("postgres.go", "CreatePostgres()", "Invalid dbURL")
		return nil
	}

	conn, err := sql.Open("postgres", dbUrl)
	logger.LogFatalError("postgres.go", "CreatePostgresConn()", "Connection error", err)

	return &Pg{
		DB: database.New(conn),
	}
}
