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
	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	sessionToken := sessionTokenCookie.Value

	filesMetadata, err := data.GetAllFilesMetadata(sessionToken)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	user, err := data.GetUserBySessionToken(sessionToken)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	data := indexPageData{TitleText: "files up n' down", Files: filesMetadata, User: *user}
	indexTemplate.Execute(w, data)
}
