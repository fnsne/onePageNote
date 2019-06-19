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
	panic("implement me")
}

func Test_Server_can_edit_note_date(t *testing.T) {
	store := &StubStore{}
	server := OnePageNoteServer{store}

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
