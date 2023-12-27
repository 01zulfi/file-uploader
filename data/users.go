package data

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/01zulfi/file-uploader/db"
)

func createToken() (string, error) {
	randBytes := make([]byte, 32)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randBytes), nil
}

type User struct {
	id       int
	username string
	password string
}

type Session struct {
	Id     int
	Token  string
	UserId int `db:"user_id"`
}

func getUser(username string) (*User, error) {
	db := db.Get()
	var user User

	err := db.QueryRow(context.Background(), "select * from users where username = $1", username).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func createSessionEntry(user User) (*Session, error) {
	token, err := createToken()
	if err != nil {
		return nil, err
	}

	db := db.Get()
	var session Session
	err = db.QueryRow(context.Background(), "insert into sessions ( token, user_id ) values ( $1, $2 ) returning *", token, user.id).Scan(&session.Id, &session.Token, &session.UserId)

	return &session, err
}

func Login(username string, password string) (*Session, error) {
	user, err := getUser(username)
	if err != nil {
		return nil, err
	}

	if user.password != password {
		return nil, err
	}

	session, err := createSessionEntry(*user)
	if err != nil {
		return nil, err
	}

	return session, nil
}
