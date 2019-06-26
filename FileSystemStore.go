package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemStore struct {
	database *os.File
	lastKey  int
}

func NewFileSystemStore(database *os.File) *FileSystemStore {
	database.Seek(0, 0)
	return &FileSystemStore{database, 0}
}

func (f *FileSystemStore) GetNote(id int) Note {
	f.database.Seek(0, 0)
	var notes []Note
	json.NewDecoder(f.database).Decode(&notes)
	fmt.Println("notes in db = ", notes)
	if id >= 1 {
		return notes[id-1]
	} else {
		return Note{}
	}
}

func (f *FileSystemStore) SetNote(id int, note Note) {
	var notes []Note
	err := json.NewDecoder(f.database).Decode(notes)
	if err != nil {
		panic(err.Error())
	}
	if id >= 1 {
		notes[id-1] = note
		json.NewEncoder(f.database).Encode(&notes)
	}
}

func (f *FileSystemStore) GetNoteList() []Note {
	var notes []Note
	json.NewDecoder(f.database).Decode(&notes)
	return notes
}

func (f *FileSystemStore) CreateNote(note Note) int {
	note.Id = f.lastKey
	f.lastKey++
	var notes []Note
	json.NewDecoder(f.database).Decode(&notes)
	notes = append(notes, note)
	json.NewEncoder(f.database).Encode(&notes)
	return note.Id
}
