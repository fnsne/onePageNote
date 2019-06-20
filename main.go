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
	notes map[int]Note
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		notes: make(map[int]Note),
	}
}

func (i *InMemoryStore) SetNote(id int, note Note) {
	i.notes[id] = note
}

func (i *InMemoryStore) GetNote(id int) Note {
	return i.notes[id]
}
