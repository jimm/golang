package main

import (
	"database/sql"
	_ "github.com/jbarham/gopgsqldriver"
	"os"
	"strings"
)

func chkerr(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// usage: go-publisher host=localhost user=user_name dbname=database_name
func main() {
	connString := strings.Join(os.Args[1:], " ")
	db, err := sql.Open("postgres", connString)
	chkerr(err)

	// Count pages
	rows, err := db.Query("SELECT COUNT(*) FROM pages")
	chkerr(err)
	if !rows.Next() {
		println("Result.Next failed")
		os.Exit(1)
	}
	var count int
	err = rows.Scan(&count)
	chkerr(err)
	println("There are", count, "pages.")
	rows.Close()

	// Retrieve inserted rows and verify inserted values.
	rows, err = db.Query("SELECT title FROM pages LIMIT 10")
	chkerr(err)
	for i := 0; rows.Next(); i++ {
		var s string

		err := rows.Scan(&s)
		if err != nil {
			println("scan error:", err)
			os.Exit(1)
		}
		println(s)
	}
	rows.Close()
}
