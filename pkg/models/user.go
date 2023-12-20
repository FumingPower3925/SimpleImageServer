package models

type User struct {
	Id       string `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
}
