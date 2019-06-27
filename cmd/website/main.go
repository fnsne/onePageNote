package main

import (
	"fmt"
	"net/http"
	main2 "onePage"
	"os"
)

const path = "notes.db.json"

func main() {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Errorf("probilem opening %s %v", path, err)
	}
	store := main2.NewFileSystemStore(db)
	server := main2.NewOnePageNoteServer(store)
	err = http.ListenAndServe(":7000", server)
	if err != nil {
		println("could not listen on port 7000")
	}
}
