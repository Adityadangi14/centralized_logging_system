package initializers

import (
	"database/sql"
	"fmt"

	db "github.com/Adityadangi14/centralized_logging_system/log_collector/db/gen"
	_ "github.com/lib/pq"
)

type DBConnector interface {
	Connect(dsn string) (*sql.DB, error)
}

type PostgresConnector struct{}

func (PostgresConnector) Connect(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

var Q *db.Queries

func InitPostgres(connector DBConnector, dsn string) (*db.Queries, error) {
	DB, err := connector.Connect(dsn)

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	queries := db.New(DB)
	Q = queries

	err = DB.Ping()

	if err != nil {
		panic("Failed to connect to db" + err.Error())
	}
	return queries, nil
}
