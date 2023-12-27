package controllers

import (
	"github.com/01zulfi/file-uploader/data"
)

type indexPageData struct {
	TitleText string
	Files     []data.FilesMetadata
}
