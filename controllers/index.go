package controllers

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	var files []indexPageFile

	rawFiles, err := os.ReadDir("./uploads")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	for _, file := range rawFiles {
		name := strings.Split(file.Name(), "___")[0]
		files = append(files, indexPageFile{Filename: name, Filelink: file.Name()})
	}

	data := indexPageData{TitleText: "files", Files: files}
	html.Execute(w, data)
}
