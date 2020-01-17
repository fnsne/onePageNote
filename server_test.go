package onePage

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type StubStore struct {
	mock.Mock
}

func (s *StubStore) GetNote(id int) Note {
	args := s.Called(id)
	return args.Get(0).(Note)
}

func (s *StubStore) SetNote(id int, note Note) {
	s.Called(id, note)
}

func (s *StubStore) GetNoteList() []Note {
	args := s.Called()
	return args.Get(0).([]Note)
}

func (s *StubStore) CreateNote(note Note) int {
	args := s.Called(note)
	return args.Int(0)
}

type ServerTests struct {
	suite.Suite
	server *OnePageNoteServer
	store  *StubStore
}

func TestSuiteInit(t *testing.T) {
	suite.Run(t, new(ServerTests))
}

func (suite *ServerTests) SetupTest() {
	suite.store = &StubStore{}
	suite.server = NewOnePageNoteServer(suite.store)
}

func (suite *ServerTests) Test_create_note() {
	note := Note{
		Id:    0,
		Date:  nil,
		Title: "new note",
		Grids: nil,
	}
	suite.serverShouldCreatedNote(note, 0)
}

func (suite *ServerTests) Test_edit_note() {
	suite.givenNoteList([]Note{{
		Id:    1,
		Date:  nil,
		Title: "new note",
		Grids: nil,
	}})
	suite.serverShouldUpdateNote(Note{
		Id:    1,
		Date:  nil,
		Title: "edited note title",
		Grids: nil,
	})
}

func (suite *ServerTests) Test_get_note_list() {
	notes := []Note{{
		Id:    1,
		Date:  nil,
		Title: "new note",
		Grids: nil,
	}}
	suite.givenNoteList(notes)
	suite.serverShouldResponseNoteList(notes)
}

func (suite *ServerTests) serverShouldResponseNoteList(notes []Note) {
	request := suite.getNoteListRequest()
	response := suite.processRequest(request)

	assert.Equal(suite.T(), http.StatusOK, response.Code, "")
	var resNotes []Note
	_ = json.NewDecoder(response.Body).Decode(&resNotes)
	assert.Equal(suite.T(), notes, resNotes)
}

func (suite *ServerTests) serverShouldUpdateNote(note Note) {
	request := suite.updateRequest(note)
	response := suite.processRequest(request)

	assert.Equal(suite.T(), http.StatusOK, response.Code, "")
	suite.store.AssertCalled(suite.T(), "SetNote", 1, note)
}

func (suite *ServerTests) willUpdateNote(note Note) *mock.Call {
	return suite.store.On("SetNote", 1, note)
}

func (suite *ServerTests) processRequest(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	suite.server.ServeHTTP(response, request)
	return response
}

func (suite *ServerTests) getNoteListRequest() *http.Request {
	request := httptest.NewRequest(http.MethodGet, "/api/note/", nil)
	return request
}

func (suite *ServerTests) updateRequest(note Note) *http.Request {
	suite.willUpdateNote(note)
	request := httptest.NewRequest(http.MethodPost, "/api/note/"+strconv.Itoa(note.Id), suite.marshal(note))
	return request
}

func (suite *ServerTests) createNoteRequest(note Note) *http.Request {
	suite.willCreateNote(note, 0)
	request := httptest.NewRequest(http.MethodPost, "/api/note/", suite.marshal(note))
	return request
}

func (suite *ServerTests) givenNoteList(notes []Note) *mock.Call {
	return suite.store.On("GetNoteList").Return(notes)
}

func (suite *ServerTests) serverShouldCreatedNote(note Note, id int) {
	request := suite.createNoteRequest(note)
	response := suite.processRequest(request)

	assert.Equal(suite.T(), http.StatusOK, response.Code, "")
	suite.assertResponseNote(response, id, note)
}

func (suite *ServerTests) assertResponseNote(response *httptest.ResponseRecorder, id int, note Note) {
	responseNote := suite.getResponseNote(response)
	assert.Equal(suite.T(), id, responseNote.Id)

	suite.store.AssertCalled(suite.T(), "CreateNote", note)
	assert.Equal(suite.T(), "new note", note.Title)
}

func (suite *ServerTests) getResponseNote(response *httptest.ResponseRecorder) Note {
	var responseNote Note
	_ = json.NewDecoder(response.Body).Decode(responseNote)
	return responseNote
}

func (suite *ServerTests) marshal(note Note) *bytes.Buffer {
	body := &bytes.Buffer{}
	_ = json.NewEncoder(body).Encode(note)
	return body
}

func (suite *ServerTests) willCreateNote(note Note, returnValue int) *mock.Call {
	return suite.store.On("CreateNote", note).Return(returnValue)
}
