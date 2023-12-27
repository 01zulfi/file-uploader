package controllers

import (
	"fmt"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Query().Get("file")
	if filepath == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	OGFilename, err := data.GetOGFilename(filepath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	contents, err := data.GetFileContents(filepath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+OGFilename)
	w.Write(contents)
}
