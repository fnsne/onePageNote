package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	date time.Time
}

func (s *StubStore) SetNote(note Note) {
	s.date = note.Date
}

func (s *StubStore) GetNote() Note {
	return Note{Date: s.date}
}

func Test_Server_can_edit_note_date(t *testing.T) {
	store := &StubStore{}
	server := NewOnePageNoteServer(store)

	body := &bytes.Buffer{}
	date, _ := time.Parse("2006-01-02", "2018-05-10")
	note := Note{Date: date}
	json.NewEncoder(body).Encode(note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, date, store.date)
}

func Test_Server_can_get_stored_note_date(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2018-05-10")
	store := &StubStore{date: date}
	server := NewOnePageNoteServer(store)

	request := httptest.NewRequest(http.MethodGet, "/api/note/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)

	var note Note
	json.NewDecoder(response.Body).Decode(&note)

	assert.Equal(t, date, note.Date)
}
