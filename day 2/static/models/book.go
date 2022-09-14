package models

type Book struct {
	Id     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Page   int    `json:"page" form:"page"`
	Author string `json:"author" form:"author"`
}
