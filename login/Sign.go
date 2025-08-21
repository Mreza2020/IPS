package Login

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Sign(name string, password string) string {
	db, err := sql.Open("mysql", "username:password@protocol(address:port)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	insertQuery := "INSERT INTO login(name, password) VALUES (?,?)"
	_, err = db.Exec(insertQuery, name, password)

	if err != nil {
		log.Fatalln(err)

	}
	return "ok"

}
