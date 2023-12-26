package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

var indexTemplate *template.Template

func init() {
	templatePath := "./templates/index.tmpl"
	indexTemplate = template.Must(template.ParseFiles(templatePath))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	filesMetadata, err := data.GetAllFilesMetadata()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	data := indexPageData{TitleText: "files up n' down", Files: filesMetadata}
	indexTemplate.Execute(w, data)
}
