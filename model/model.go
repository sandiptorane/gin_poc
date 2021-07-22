package model

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
