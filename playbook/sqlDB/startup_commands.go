package sqlDB

import (
	"database/sql"
	"log"
"fmt"
	_ "github.com/mattn/go-sqlite3"
 "path/filepath"

)

func SQLStartUp() *sql.DB {

 absPath, err := filepath.Abs("./mydatabase.db")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(absPath)
	db, err := sql.Open("sqlite3", "./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS userinfo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	fmt.Println(db)

	return db
}
