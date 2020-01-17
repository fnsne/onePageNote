package onePage

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Store interface {
	GetNote(id int) Note
	SetNote(id int, note Note)
	GetNoteList() []Note //should sorted by created date
	CreateNote(note Note) int
}

type Note struct {
	Id    int
	Date  *time.Time
	Title string
	Grids []Grid
}

type NoteServer struct {
	store Store
	http.Handler
}

func NewOnePageNoteServer(store Store) *NoteServer {
	server := new(NoteServer)

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.homePage))
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.Handle("/api/note/", http.HandlerFunc(server.note))

	server.Handler = router
	server.store = store
	return server
}

func (s *NoteServer) homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp, _ := template.ParseFiles("view/note.html")
		notes := s.store.GetNoteList()
		currentNoteId := 1
		if len(notes) != 0 {
			currentNoteId = notes[len(notes)-1].Id
		}
		_ = tmp.Execute(w, currentNoteId)
	}
}

func (s *NoteServer) updateNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getNoteId(r))
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	if len(s.store.GetNoteList()) == 0 {
		s.store.CreateNote(note)
	}
	s.store.SetNote(id, note)
	w.WriteHeader(http.StatusOK)
}

func (s *NoteServer) retrieveNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getNoteId(r))
	err := json.NewEncoder(w).Encode(s.store.GetNote(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (s *NoteServer) listNote(w http.ResponseWriter) {
	var notes = s.store.GetNoteList()
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (s *NoteServer) createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	id := s.store.CreateNote(note)
	w.WriteHeader(http.StatusOK)
	note.Id = id
	err := json.NewEncoder(w).Encode(&note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}

func (s *NoteServer) note(w http.ResponseWriter, r *http.Request) {
	if requestHasId(r) {
		if r.Method == http.MethodGet {
			confirmId(s.retrieveNote)(w, r)
		}
		if r.Method == http.MethodPost {
			confirmId(confirmInputNote(s.updateNote))(w, r)
		}
	} else {
		if r.Method == http.MethodGet {
			s.listNote(w)
		}
		if r.Method == http.MethodPost {
			confirmInputNote(s.createNote)(w, r)
		}
	}
}

func requestHasId(r *http.Request) bool {
	idString := getNoteId(r)
	hasId := len(idString) != 0
	return hasId
}

func getNoteId(r *http.Request) string {
	idString := r.URL.Path[len("/api/note/"):]
	return idString
}

type Grid struct {
	Keyword string
	Comment string
}
