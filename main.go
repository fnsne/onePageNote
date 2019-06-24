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

func (i *InMemoryStore) CreateNote(note Note) {
	panic("implement me")
}

func (i *InMemoryStore) GetNoteList() []Note {
	var notes []Note
	for _, note := range i.notes {
		notes = append(notes, note)
	}
	return notes
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
