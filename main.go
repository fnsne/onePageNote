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
	notes   []Note
	lastKey int
}

func (i *InMemoryStore) CreateNote(note Note) int {
	note.Id = i.lastKey
	i.notes = append(i.notes, note)
	i.lastKey++
	return note.Id
}

func (m *InMemoryStore) GetNoteList() []Note {
	var notes []Note
	for i := len(m.notes) - 1; i >= 0; i-- {
		notes = append(notes, m.notes[i])
	}
	return notes
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		lastKey: 1,
	}
}

func (i *InMemoryStore) SetNote(id int, note Note) {
	if id >= 1 {
		i.notes[id-1] = note
	}
}

func (i *InMemoryStore) GetNote(id int) Note {
	if id >= 1 && len(i.notes) >= id {
		return i.notes[id-1]
	}
	return Note{}
}
