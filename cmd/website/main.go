package main

import (
	"net/http"
	"onePage"
	"os"
)

const path = "notes.db.json"

func main() {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("problem opening " + path + " " + err.Error())
	}
	store := onePage.NewFileSystemStore(db)
	server := onePage.NewOnePageNoteServer(store)
	err = http.ListenAndServe(":7000", server)
	if err != nil {
		println("could not listen on port 7000")
	}
}
