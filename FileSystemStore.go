package main

import (
	"encoding/json"
	"os"
)

type FileSystemStore struct {
	database *os.File
	lastKey  int
	notes    []Note
}

func NewFileSystemStore(database *os.File) *FileSystemStore {
	var notes []Note
	database.Seek(0, 0)
	json.NewDecoder(database).Decode(&notes)
	lastKey := 1
	if len(notes) != 0 {
		lastKey = notes[len(notes)-1].Id
	}
	return &FileSystemStore{database, lastKey, notes}
}

func (f *FileSystemStore) GetNote(id int) Note {
	if id >= 1 && len(f.notes) >= id {
		return f.notes[id-1]
	} else {
		return Note{}
	}
}

func (f *FileSystemStore) SetNote(id int, note Note) {
	if id >= 1 {
		f.notes[id-1] = note
		f.database.Seek(0, 0)
		json.NewEncoder(f.database).Encode(&f.notes)
	}
}

func (f *FileSystemStore) GetNoteList() []Note {
	var notes []Note
	for i := len(f.notes) - 1; i >= 0; i-- {
		notes = append(notes, f.notes[i])
	}
	return notes
}

func (f *FileSystemStore) CreateNote(note Note) int {
	note.Id = f.lastKey
	f.lastKey++
	f.notes = append(f.notes, note)
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(&f.notes)
	return note.Id
}
