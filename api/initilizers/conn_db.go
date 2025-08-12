package initilizers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/Adityadangi14/centralized_logging_system/api/db/gen"
	_ "github.com/lib/pq"
)

var Q *db.Queries

func ConnectToPostgres() {
	dsn := os.Getenv("postgresurl")
	fmt.Println("dsn", dsn)
	var err error
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("❌ Unable to connect to database: %v\n", err)
	}

	fmt.Println("Connected to Supabase Postgres")

	Q = db.New(DB)
}
