package Login

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Login(username string, password string) string {
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

	var name string

	err = db.QueryRow("SELECT name FROM login WHERE password = ? ", password).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return "err"

	} else if err != nil {
		return "err"

	} else {
		if name == username {
			return "ok"
		} else {
			return "err"
		}
	}

}
