package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	date  time.Time
	Title string
	note  Note
}

func (s *StubStore) SetNote(note Note) {
	s.note = note
}

func (s *StubStore) GetNote() Note {
	return s.note
}

func Test_Server_can_store_note_title(t *testing.T) {
	store := &StubStore{}
	server := NewOnePageNoteServer(store)

	wantedTitle := "我是主題"
	note := Note{Title: wantedTitle}
	body := createNoteJSONBody(t, note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, wantedTitle, store.GetNote().Title)
}

func Test_Server_can_edit_note_date(t *testing.T) {
	store := &StubStore{}
	server := NewOnePageNoteServer(store)

	date, _ := time.Parse("2006-01-02", "2018-05-10")
	note := Note{Date: &date}
	body := createNoteJSONBody(t, note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, date, *store.GetNote().Date)
}

func Test_Server_can_get_stored_note_date(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2018-05-10")
	store := &StubStore{note: Note{Date: &date}}
	server := NewOnePageNoteServer(store)

	request := httptest.NewRequest(http.MethodGet, "/api/note/", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
	note := getResponseNote(t, response)
	assert.Equal(t, &date, note.Date)
}

func getResponseNote(t *testing.T, response *httptest.ResponseRecorder) Note {
	var note Note
	err := json.NewDecoder(response.Body).Decode(&note)
	assert.NoError(t, err)
	return note
}

func createNoteJSONBody(t *testing.T, note Note) io.Reader {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(note)
	assert.NoError(t, err)
	return body
}
