package controllers

import (
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	sessionToken := sessionTokenCookie.Value

	_ = data.DeleteSession(sessionToken)

	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/"})
	http.Redirect(w, r, "/login", http.StatusFound)
}
