package main

import (
	"encoding/json"
	"os"
)

type FileSystemStore struct {
	database *os.File
	lastKey  int
	notes   []Note
}

func NewFileSystemStore(database *os.File) *FileSystemStore {
	database.Seek(0, 0)
	var notes []Note
	json.NewDecoder(database).Decode(&notes)
	return &FileSystemStore{database, 0, notes}
}

func (f *FileSystemStore) GetNote(id int) Note {
	if id >= 1 {
		return f.notes[id-1]
	} else {
		return Note{}
	}
}

func (f *FileSystemStore) SetNote(id int, note Note) {
	if id >= 1 {
		f.notes[id-1] = note
		f.database.Seek(0,0)
		json.NewEncoder(f.database).Encode(&f.notes)
	}
}

func (f *FileSystemStore) GetNoteList() []Note {
	return f.notes
}

func (f *FileSystemStore) CreateNote(note Note) int {
	note.Id = f.lastKey
	f.lastKey++
	f.notes = append(f.notes, note)
	json.NewEncoder(f.database).Encode(&f.notes)
	return note.Id
}
