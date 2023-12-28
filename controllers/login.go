package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

var t *template.Template

func init() {
	templatePath := "./templates/login.tmpl"
	t = template.Must(template.ParseFiles(templatePath))
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		loginGet(w, r)
	} else {
		loginPost(w, r)
	}
}

func loginGet(w http.ResponseWriter, r *http.Request) {
	t.Execute(w, nil)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	session, err := data.Login(string(username), string(password))
	if ive, ok := err.(*data.InvalidPasswordError); ok {
		http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/"})
		t.Execute(w, struct {
			Error    string
			Username string
		}{Error: ive.Error(), Username: string(username)})
		return
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: session.Token, Path: "/"})
	http.Redirect(w, r, "/", http.StatusFound)
}
