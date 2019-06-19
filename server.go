package main

import (
	"html/template"
	"net/http"
)

type OnePageNoteServer struct {
}

func (*OnePageNoteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmp, _ := template.ParseFiles("view/note.html")
		if r.Method == http.MethodGet {
			tmp.Execute(w, nil)
		}
	}))
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.ServeHTTP(w, r)
}
