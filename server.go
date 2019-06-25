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
	GetNoteList() []Note
	CreateNote(note Note)
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
func (receiver OnePageNoteServer) homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmp, _ := template.ParseFiles("view/note.html")
		currentNoteId := 1
		tmp.Execute(w, currentNoteId)
	}
}
func (s *OnePageNoteServer) note(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idString := r.URL.Path[len("/api/note/"):]
		if len(idString) == 0 {
			var note Note
			err := json.NewDecoder(r.Body).Decode(&note)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			s.store.CreateNote(note)
			w.WriteHeader(http.StatusOK)
		} else {
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
			} else {
				if len(s.store.GetNoteList()) == 0 {
					s.store.CreateNote(note)
				}
				s.store.SetNote(id, note)
				//fmt.Println("set note ", id, "to ", note)
				w.WriteHeader(http.StatusOK)
			}
			return
		}
	}
	if r.Method == http.MethodGet {
		idString := r.URL.Path[len("/api/note/"):]
		if len(idString) == 0 {
			var notes = s.store.GetNoteList()
			err := json.NewEncoder(w).Encode(notes)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				fmt.Println("get note list", notes)
			}
		} else {
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
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

type Grid struct {
	Keyword string
	Comment string
}
