package onePage

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func copyRequestBody(r *http.Request) (copiedBody io.ReadCloser) {
	buf, _ := ioutil.ReadAll(r.Body)
	temp := ioutil.NopCloser(bytes.NewBuffer(buf))
	copiedBody = ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = temp
	return copiedBody
}

func confirmId(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := getNoteId(r)
		_, err := strconv.Atoi(idString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			f(w, r)
		}
	}
}

func confirmInputNote(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note Note
		body := copyRequestBody(r)
		err := json.NewDecoder(body).Decode(&note)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			f(w, r)
		}
	}
}
