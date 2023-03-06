package initializers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectDatabase() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, os.Getenv("DB_PW"), os.Getenv("DB_NAME"))

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}
