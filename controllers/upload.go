package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["files"]
	var filenames []string
	var buffer bytes.Buffer
	for _, fh := range fhs {
		f, err := fh.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.Copy(&buffer, f)
		formattedFilename := fh.Filename + "___" + time.Now().String()
		os.WriteFile("./uploads/"+formattedFilename, buffer.Bytes(), 0666)
		buffer.Reset()
		filenames = append(filenames, fh.Filename)
	}

	html, err := template.ParseFiles("./templates/upload.tmpl")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	var data []uploadPageData
	for _, filename := range filenames {
		data = append(data, uploadPageData{Filename: filename})
	}
	html.Execute(w, data)
}
