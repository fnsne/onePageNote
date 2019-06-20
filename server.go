package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Store interface {
	GetNote() Note
	SetNote(note Note)
}

type Note struct {
	Date *time.Time
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
func (s *OnePageNoteServer) note(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var note Note
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			s.store.SetNote(note)
			w.WriteHeader(http.StatusOK)
		}
	} else if r.Method == http.MethodGet {
		err := json.NewEncoder(w).Encode(s.store.GetNote())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *OnePageNoteServer) homePage(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("view/note.html")
	if r.Method == http.MethodGet {
		tmp.Execute(w, nil)
	}
}
