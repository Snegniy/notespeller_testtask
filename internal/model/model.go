package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type Note struct {
	Id         int    `json:"id" db:"id"`
	UserId     int    `json:"userid" db:"userid"`
	DateCreate string `json:"date_create" db:"date_create"`
	Note       string `json:"note" db:"note"`
}
