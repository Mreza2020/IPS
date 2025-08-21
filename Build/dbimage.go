package Build

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DbImage(adder string, name string) bool {
	db, err := sql.Open("mysql", "username:password@protocol(address:port)/dbname")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)

		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println("Connected to MySQL!")

	insertQuery := "INSERT INTO image(name, adder) VALUES (?,?)"
	_, err = db.Exec(insertQuery, name, adder)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true

}
