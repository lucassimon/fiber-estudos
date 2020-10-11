package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// CreateConnection create connection with postgres db
func CreateConnection() *sql.DB {
	var db *sql.DB
	// load .env file
	// err := godotenv.Load("../../.env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	fmt.Println(os.Getenv("POSTGRES_URL"))

	// Open the connection
	db, err := sql.Open("postgres", "postgresql://estudos:teste123@localhost:25432/fiberestudos?sslmode=disable")

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
