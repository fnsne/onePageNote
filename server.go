package main

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
}

type Note struct {
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
	router.Handle("/note/", http.HandlerFunc(server.notePage))
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.Handle("/api/note/", http.HandlerFunc(server.note))

	server.Handler = router
	server.store = store
	return server
}
func (receiver OnePageNoteServer) homePage(w http.ResponseWriter, r *http.Request) {

}
func (s *OnePageNoteServer) note(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.URL.Path[len("/api/note/"):])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("err = ", err)
			return
		}
		var note Note
		err = json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			s.store.SetNote(id, note)
			fmt.Println("set note ", note)
			w.WriteHeader(http.StatusOK)
		}
		return
	}
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.URL.Path[len("/api/note/"):])
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
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (s *OnePageNoteServer) notePage(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("view/note.html")
	if r.Method == http.MethodGet {
		tmp.Execute(w, nil)
	}
}

type Grid struct {
	Keyword string
	Comment string
}
