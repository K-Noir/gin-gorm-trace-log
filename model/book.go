package model

type Book struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Auther string `json:"auther,omitempty"`
}
