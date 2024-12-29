package models

type Users struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
