package controllers

import (
	"net/http"
	"os"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	contents, err := os.ReadFile("./uploads/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(contents)
}
