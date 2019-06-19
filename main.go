package main

import (
	"net/http"
)

func main() {
	server := &OnePageNoteServer{}
	err := http.ListenAndServe(":7000", server)
	if err != nil {
		println("could not listen on port 7000")
	}
}