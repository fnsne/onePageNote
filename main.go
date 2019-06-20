package main

import (
	"net/http"
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
	return &InMemoryStore{
		note: Note{},
	}
}

func (i *InMemoryStore) SetNote(note Note) {
	i.note = note
}

func (i *InMemoryStore) GetNote() Note {
	return i.note
}
