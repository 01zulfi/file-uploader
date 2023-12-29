package data

import (
	"context"
	"io/fs"
	"os"
	"time"

	"github.com/01zulfi/file-uploader/db"
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
	var filesMetadata []FilesMetadata
	db := db.Get()
	rows, err := db.Query(context.Background(), "select * from files")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row struct {
			id       int
			filename string
			filepath string
			owner    int
		}
		err := rows.Scan(&row.id, &row.filename, &row.filepath, &row.owner)
		if err != nil {
			return nil, err
		}
		filesMetadata = append(filesMetadata, FilesMetadata{OGFilename: row.filename, Filepath: row.filepath})
	}

	return filesMetadata, nil
}

func SaveFile(filename string, contents []byte) (FilesMetadata, error) {
	formattedFilename := time.Now().String() + separator + filename
	err := os.WriteFile(uploadDir+forwardSlash+formattedFilename, contents, filePerm)
	if err != nil {
		return FilesMetadata{}, err
	}

	db := db.Get()
	_, err = db.Exec(context.Background(), "insert into files (filename, filepath, owner) values ($1, $2, $3)", filename, formattedFilename, 1)
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

func GetOGFilename(filepath string) (string, error) {
	db := db.Get()
	var filename string
	err := db.QueryRow(context.Background(), "select filename from files where filepath = $1", filepath).Scan(&filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func DeleteFile(filepath string) error {
	db := db.Get()
	_, err := db.Exec(context.Background(), "delete from files where filepath = $1", filepath)
	return err
}
