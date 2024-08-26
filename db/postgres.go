package db

import (
	// "database/sql"
	appConfig "foo/api/config" 
	"os"

	_ "github.com/lib/pq"
)

func setupDb(host string) {
	var pgPassword = os.Getenv("pgPassword")
	var dbConfig = appConfig.PgConfig(
		""
	) 
}

// func main() {
// 	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	age := 21
// 	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
// }
