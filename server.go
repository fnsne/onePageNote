package onePage

import (
	"html/template"
	"net/http"
)

type OnePageNoteServer struct {
}

func (*OnePageNoteServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("note.html")
	if r.Method == http.MethodGet {
		tmp.Execute(w, nil)
	}
}
