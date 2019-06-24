package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	notes map[int]Note
}

func (s *StubStore) CreateNote(note Note) {
	s.notes[5] = note
}

func (s *StubStore) GetNoteList() []Note {
	var notes []Note
	for _, note := range s.notes {
		notes = append(notes, note)
	}
	return notes
}

func NewStubStore() *StubStore {
	return &StubStore{notes: make(map[int]Note)}
}

func (s *StubStore) SetNote(id int, note Note) {
	s.notes[id] = note
}

func (s *StubStore) GetNote(id int) Note {
	return s.notes[id]
}

func Test_Server_can_create_new_note(t *testing.T) {
	store := &StubStore{notes: map[int]Note{1: {Id: 1, Title: "title1"}, 2: {Id: 2, Title: "title2"}}}
	server := NewOnePageNoteServer(store)

	response := httptest.NewRecorder()
	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(Note{Title: "new note"})
	request := httptest.NewRequest(http.MethodPost, "/api/note/", body)

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, 3, len(store.GetNoteList()))
}

func Test_Server_can_get_list_of_notes(t *testing.T) {
	store := &StubStore{notes: map[int]Note{1: {Id: 1, Title: "title1"}, 2: {Id: 2, Title: "title2"}}}
	server := NewOnePageNoteServer(store)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/note/", nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	var notes []Note
	fmt.Println(store.GetNoteList())
	json.NewDecoder(response.Body).Decode(&notes)
	assert.Equal(t, 1, notes[0].Id)
	assert.Equal(t, 2, notes[1].Id)
	assert.Equal(t, "title1", notes[0].Title)
	assert.Equal(t, "title2", notes[1].Title)
}

func Test_Server_can_store_other_note(t *testing.T) {
	store := &StubStore{notes: map[int]Note{1: {Title: "title1"}, 2: {Title: "title2"}}}
	server := NewOnePageNoteServer(store)

	response := httptest.NewRecorder()
	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(Note{Title: "titleA"})
	request := httptest.NewRequest(http.MethodPost, "/api/note/2", body)

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "titleA", store.GetNote(2).Title)
}

func Test_Server_can_get_other_note(t *testing.T) {
	store := &StubStore{notes: map[int]Note{1: {Title: "title1"}, 2: {Title: "title2"}}}
	server := NewOnePageNoteServer(store)

	response := httptest.NewRecorder()
	request1 := httptest.NewRequest(http.MethodGet, "/api/note/1", nil)
	request2 := httptest.NewRequest(http.MethodGet, "/api/note/2", nil)

	server.ServeHTTP(response, request1)

	assert.Equal(t, http.StatusOK, response.Code)
	var note1 Note
	json.NewDecoder(response.Body).Decode(&note1)
	assert.Equal(t, "title1", note1.Title)

	server.ServeHTTP(response, request2)

	assert.Equal(t, http.StatusOK, response.Code)
	var note2 Note
	json.NewDecoder(response.Body).Decode(&note2)
	assert.Equal(t, "title2", note2.Title)
}

func Test_Server_can_store_keyword_and_comment(t *testing.T) {
	store := NewStubStore()
	server := NewOnePageNoteServer(store)

	wantedKeyword := "關鍵字1"
	wantedComment := "評論1"
	grids := []Grid{{Keyword: wantedKeyword, Comment: wantedComment}}
	note := Note{Grids: grids}
	body := createNoteJSONBody(t, note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/1", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, wantedKeyword, store.GetNote(1).Grids[0].Keyword)
	assert.Equal(t, wantedComment, store.GetNote(1).Grids[0].Comment)
}

func Test_Server_can_store_note_title(t *testing.T) {
	store := NewStubStore()
	server := NewOnePageNoteServer(store)

	wantedTitle := "我是主題"
	note := Note{Title: wantedTitle}
	body := createNoteJSONBody(t, note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/1", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, wantedTitle, store.GetNote(1).Title)
}

func Test_Server_can_edit_note_date(t *testing.T) {
	store := NewStubStore()
	server := NewOnePageNoteServer(store)

	date, _ := time.Parse("2006-01-02", "2018-05-10")
	note := Note{Date: &date}
	body := createNoteJSONBody(t, note)

	request := httptest.NewRequest(http.MethodPost, "/api/note/1", body)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, date, *store.GetNote(1).Date)
}

func Test_Server_can_get_stored_note_date(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2018-05-10")
	store := &StubStore{notes: map[int]Note{1: {Date: &date}}}
	server := NewOnePageNoteServer(store)

	request := httptest.NewRequest(http.MethodGet, "/api/note/1", nil)
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
