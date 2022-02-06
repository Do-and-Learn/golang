package main

import (
	"os"
	"text/template"
)

type Entry struct {
	Number int
	Square int
}

func main() {
	tFile := "text.gotext"
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var Entries []Entry
	for _, i := range DATA {
		entry := Entry{i[0], i[1]}
		Entries = append(Entries, entry)
	}
	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
