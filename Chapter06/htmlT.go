package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type Entry struct {
	Number int
	Double int
	Square int
}

func myHandler(data []Entry) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		const tFile string = "html.gohtml"
		tmp := template.Must(template.ParseGlob(tFile))
		tmp.ExecuteTemplate(writer, tFile, data)
	}
}

func main() {
	_, err := sql.Open("sqlite3", "htmlT.db")

	var data []Entry
	for i := 20; i < 50; i++ {
		data = append(data, Entry{i, i * 2, i * i})
	}

	http.HandleFunc("/", myHandler(data))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
