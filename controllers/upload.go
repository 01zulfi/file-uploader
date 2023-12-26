package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

var uploadTemplate *template.Template

func init() {
	templatePath := "./templates/upload.tmpl"
	uploadTemplate = template.Must(template.ParseFiles(templatePath))
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["files"]
	var filesMetadata []data.FilesMetadata
	var buffer bytes.Buffer
	for _, fh := range fhs {
		f, err := fh.Open()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.Copy(&buffer, f)
		fileMetadata, err := data.SaveFile(fh.Filename, buffer.Bytes())
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		buffer.Reset()
		filesMetadata = append(filesMetadata, fileMetadata)
	}

	uploadTemplate.Execute(w, filesMetadata)
}
