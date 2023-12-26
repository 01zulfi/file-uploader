package controllers

import (
	"net/http"

	"github.com/01zulfi/file-uploader/data"
)

type handler func(w http.ResponseWriter, r *http.Request)

type indexPageData struct {
	TitleText string
	Files     []data.FilesMetadata
}
