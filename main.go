package main

import (
	"net/http"
	"time"
)

func main() {
	store := NewInMemoryStore()
	server := NewOnePageNoteServer(store)
	err := http.ListenAndServe(":7000", server)
	if err != nil {
		println("could not listen on port 7000")
	}
}

type InMemoryStore struct {
	note Note
}

func NewInMemoryStore() *InMemoryStore {
	d, _ := time.Parse("2006-01-02", "1997-11-11")
	return &InMemoryStore{
		note: Note{&d},
	}
}

func (i *InMemoryStore) SetNote(note Note) {
	i.note = note
}

func (i *InMemoryStore) GetNote() Note {

	return i.note
}
