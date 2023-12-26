package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

const port = "8080"

type IndexPageData struct {
	TitleText string
	Filenames []UploadPageData
}

type UploadPageData struct {
	Filename string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	var filenames []UploadPageData

	files, err := os.ReadDir("./uploads")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	for _, file := range files {
		filename := UploadPageData{Filename: file.Name()}
		filenames = append(filenames, filename)
	}

	data := IndexPageData{TitleText: "files", Filenames: filenames}
	html.Execute(w, data)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
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
		os.WriteFile("./uploads/"+fh.Filename, buffer.Bytes(), 0666)
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

	var data []UploadPageData
	for _, filename := range filenames {
		data = append(data, UploadPageData{Filename: filename})
	}
	html.Execute(w, data)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
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

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", handleIndex)
	server.HandleFunc("/upload", handleUpload)
	server.HandleFunc("/download/", handleDownload)

	err := http.ListenAndServe(":"+port, server)

	if err != nil {
		fmt.Println("error while starting server")
	}
}
