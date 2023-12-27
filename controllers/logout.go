package controllers

import (
	"fmt"
	"net/http"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout")
	http.Redirect(w, r, "/login", http.StatusFound)
}
