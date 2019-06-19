package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

type Store interface {
	GetNote() Note
	SetNote(note Note)
}

type Note struct {
	Date time.Time
}
type OnePageNoteServer struct {
	store Store
}

func (s *OnePageNoteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, _ := template.ParseFiles("view/note.html")
		if r.Method == http.MethodGet {
			tmp.Execute(w, nil)
		}
	}))
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.Handle("/api/note/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var note Note
			json.NewDecoder(r.Body).Decode(&note)
			s.store.SetNote(note)
			w.WriteHeader(http.StatusOK)
		}
		w.WriteHeader(http.StatusBadRequest)
	}))
	router.ServeHTTP(w, r)
}
