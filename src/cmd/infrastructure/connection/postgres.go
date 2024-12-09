package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=ms-user sslmode=disable")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}

func TestConnectPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=ms-user sslmode=disable host=localhost port=5435")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}
