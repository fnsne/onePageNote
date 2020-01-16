package onePage

import (
	"encoding/json"
	"fmt"
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
type OnePageNoteServer struct {
	store Store
	http.Handler
}

func NewOnePageNoteServer(store Store) *OnePageNoteServer {
	server := new(OnePageNoteServer)

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.homePage))
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.Handle("/api/note/", http.HandlerFunc(server.note))

	server.Handler = router
	server.store = store
	return server
}
func (s *OnePageNoteServer) homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp, _ := template.ParseFiles("view/note.html")
		notes := s.store.GetNoteList()
		currentNoteId := 1
		if len(notes) != 0 {
			currentNoteId = notes[len(notes)-1].Id
		}
		tmp.Execute(w, currentNoteId)
	}
}
func (s *OnePageNoteServer) createNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := s.store.CreateNote(note)
	w.WriteHeader(http.StatusOK)
	note.Id = id
	err = json.NewEncoder(w).Encode(&note)
	if err != nil {
		fmt.Println("err in post to create note return id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}
func (s *OnePageNoteServer) updateNote(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/api/note/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("err = ", err)
		return
	}
	var note Note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var p []byte
		r.Body.Read(p)
		fmt.Println("r.Body = ", p)
		fmt.Println("err = ", err)
	} else {
		if len(s.store.GetNoteList()) == 0 {
			s.store.CreateNote(note)
		}
		s.store.SetNote(id, note)
		fmt.Println("set note ", id, "to ", note)
		w.WriteHeader(http.StatusOK)
	}
	return
}
func (s *OnePageNoteServer) retrieveNote(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/api/note/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("err= ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(s.store.GetNote(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}
func (s *OnePageNoteServer) listNote(w http.ResponseWriter, r *http.Request) {
	var notes = s.store.GetNoteList()
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println("get note list", notes)
	}
	return
}
func (s *OnePageNoteServer) note(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idString := r.URL.Path[len("/api/note/"):]
		if len(idString) == 0 {
			fmt.Println("create note")
			s.createNote(w, r)
		} else {
			fmt.Println("update note")
			s.updateNote(w, r)
		}
	} else if r.Method == http.MethodGet {
		idString := r.URL.Path[len("/api/note/"):]
		if len(idString) == 0 {
			//fmt.Println("list note")
			s.listNote(w, r)
		} else {
			fmt.Println("retrieve note")
			s.retrieveNote(w, r)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

type Grid struct {
	Keyword string
	Comment string
}
