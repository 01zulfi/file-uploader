package controllers

import (
	"fmt"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	contents, err := data.GetFileContents(filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Write(contents)
}
