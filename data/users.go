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

func userExists(username string) (bool, error) {
	db := db.Get()
	var count int
	err := db.QueryRow(context.Background(), "select count(*) from users where username = $1", username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func getUser(username string) (*User, error) {
	db := db.Get()
	var user User

	err := db.QueryRow(context.Background(), "select * from users where username = $1", username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func createUser(username string, password string) (*User, error) {
	db := db.Get()
	var user User

	err := db.QueryRow(context.Background(), "insert into users ( username, password ) values ( $1, $2 ) returning *", username, password).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func createSessionEntry(userId int) (*Session, error) {
	token, err := createToken()
	if err != nil {
		return nil, err
	}

	db := db.Get()
	var session Session
	err = db.QueryRow(context.Background(), "insert into sessions ( token, user_id ) values ( $1, $2 ) returning *", token, userId).Scan(&session.Id, &session.Token, &session.UserId)

	return &session, err
}

func Login(username string, password string) (*Session, error) {
	var user *User
	var err error

	if u, _ := userExists(username); u {
		user, err = getUser(username)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = createUser(username, password)
		if err != nil {
			return nil, err
		}
	}

	if user.Password != password {
		return nil, &InvalidPasswordError{message: "invalid password"}
	}

	session, err := createSessionEntry(user.Id)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func Authenticate(sessionToken string) (bool, error) {
	db := db.Get()
	var session Session

	err := db.QueryRow(context.Background(), "select * from sessions where token = $1", sessionToken).Scan(&session.Id, &session.Token, &session.UserId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteSession(sessionToken string) error {
	db := db.Get()
	_, err := db.Exec(context.Background(), "delete from sessions where token = $1", sessionToken)

	return err
}
