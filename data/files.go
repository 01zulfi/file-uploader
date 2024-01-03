package data

import (
	"context"
	"fmt"
	"time"

	"github.com/01zulfi/file-uploader/db"
)

var separator = "__ID__"

type FilesMetadata struct {
	OGFilename string
	Filepath   string
}

func GetAllFilesMetadata(sessionToken string) ([]FilesMetadata, error) {
	var filesMetadata []FilesMetadata
	db := db.Get()
	rows, err := db.Query(context.Background(), `
	SELECT *
	FROM files
	WHERE owner = (
		SELECT user_id as id
		FROM sessions
		WHERE token = $1
	)	
	`, sessionToken)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var row File
		err := rows.Scan(&row.Id, &row.Filename, &row.UniqueFilename, nil, nil)
		if err != nil {
			return nil, err
		}
		filesMetadata = append(filesMetadata, FilesMetadata{OGFilename: row.Filename, Filepath: row.UniqueFilename})
	}

	return filesMetadata, nil
}

func SaveFile(filename string, contents []byte, sessionToken string) (FilesMetadata, error) {

	uniqueId, err := createToken()
	if err != nil {
		uniqueId = time.Now().String()
	}

	uniqueFilename := uniqueId + separator + filename

	db := db.Get()
	_, err = db.Exec(context.Background(), `insert into files (unique_filename, filename, content, owner) values ($1, $2, $3, (
		SELECT user_id as id
		FROM sessions
		WHERE token = $4
	)) returning id`, uniqueFilename, filename, contents, sessionToken)

	if err != nil {
		return FilesMetadata{}, err
	}

	return FilesMetadata{OGFilename: filename, Filepath: uniqueFilename}, nil
}

func GetFileContents(uniqueFilename string) ([]byte, error) {
	db := db.Get()
	var contents []byte
	err := db.QueryRow(context.Background(), "select content from files where unique_filename = $1", uniqueFilename).Scan(&contents)

	if err != nil {
		return nil, err
	}
	return contents, nil
}

func GetOGFilename(uniqueFilename string) (string, error) {
	db := db.Get()
	var filename string
	err := db.QueryRow(context.Background(), "select filename from files where unique_filename = $1", uniqueFilename).Scan(&filename)
	if err != nil {
		fmt.Println("here")
		return "", err
	}
	return filename, nil
}

func DeleteFile(filepath string) error {
	db := db.Get()
	_, err := db.Exec(context.Background(), "delete from files where filepath = $1", filepath)
	return err
}
