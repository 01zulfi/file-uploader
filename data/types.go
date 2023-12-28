package data

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Session struct {
	Id     int    `db:"id"`
	Token  string `db:"token"`
	UserId int    `db:"user_id"`
}

type File struct {
	Id       int    `db:"id"`
	Filename string `db:"filename"`
	Filepath string `db:"filepath"`
	Owner    int    `db:"owner"`
}

type InvalidPasswordError struct {
	message string
}

func (e *InvalidPasswordError) Error() string {
	return e.message
}
