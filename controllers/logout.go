package controllers

import (
	"net/http"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Path: "/"})
	http.Redirect(w, r, "/login", http.StatusFound)
}
