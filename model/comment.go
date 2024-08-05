package model

type Comment struct {
	Id      int    `json:"id"`
	BookId  int    `json:"book_id"`
	Content string `json:"content"`
}
