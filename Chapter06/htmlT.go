package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

func myHandler(data []Entry) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Host: %s Path: %s\n", request.Host, request.URL.Path)
		const tFile string = "html.gohtml"
		tmp := template.Must(template.ParseGlob(tFile))
		tmp.ExecuteTemplate(writer, tFile, data)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "htmlT.db")
	if err != nil {
		log.Panicln(err)
	}

	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		log.Panicln(err)
	}

	stmt, _ := db.Prepare("INSERT INTO data (number, double, square) VALUES (?, ?, ?)")
	for i := 20; i < 50; i++ {
		stmt.Exec(i, i*2, i*i)
	}

	row, err := db.Query("SELECT * FROM data")
	if err != nil {
		log.Panicln(err)
	}

	var data []Entry
	var n, d, s int
	for row.Next() {
		row.Scan(&n, &d, &s)
		data = append(data, Entry{n, d, s})
	}

	http.HandleFunc("/", myHandler(data))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
