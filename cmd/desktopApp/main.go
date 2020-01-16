package main

import (
	"github.com/zserge/lorca"
	"net/http"
	"onePage"
	"os"
)

const path = "notes.db.json"

func main() {
	store, err := FileStore()

	go runServer(store, "7000")

	ui, err := lorca.New("", "", 1024, 700)
	if err != nil {
		panic(err)
	}
	_ = ui.Load("localhost:7000")

	<-ui.Done()
}

func runServer(store onePage.Store, port string) {
	server := onePage.NewOnePageNoteServer(store)
	err := http.ListenAndServe(":"+port, server)
	if err != nil {
		println("could not listen on port 7000")
	}
}

func FileStore() (*onePage.FileSystemStore, error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic("problem opening " + path + " " + err.Error())
	}
	store := onePage.NewFileSystemStore(db)
	return store, err
}
