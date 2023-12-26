package data

import (
	"io/fs"
	"os"
	"strings"
	"time"
)

var (
	separator                = "__ID__"
	uploadDir                = "./uploads"
	forwardSlash             = "/"
	filePerm     fs.FileMode = 0666
)

type FilesMetadata struct {
	OGFilename string
	Filepath   string
}

func GetAllFilesMetadata() ([]FilesMetadata, error) {
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		return nil, err
	}

	var filesMetadata []FilesMetadata

	for _, file := range files {
		name := strings.Split(file.Name(), separator)[1]
		filesMetadata = append(filesMetadata, FilesMetadata{OGFilename: name, Filepath: file.Name()})
	}

	return filesMetadata, nil
}

func SaveFile(filename string, contents []byte) (FilesMetadata, error) {
	formattedFilename := time.Now().String() + separator + filename
	err := os.WriteFile(uploadDir+forwardSlash+formattedFilename, contents, filePerm)
	if err != nil {
		return FilesMetadata{}, err
	}
	return FilesMetadata{OGFilename: filename, Filepath: formattedFilename}, nil
}

func GetFileContents(filepath string) ([]byte, error) {
	contents, err := os.ReadFile(uploadDir + forwardSlash + filepath)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
