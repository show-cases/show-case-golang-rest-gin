package model

type Book struct {
	Id           int    `json:"id"`
	Csbn         string `json:"csbn"`
	Title        string `json:"title"`
	AuthorId     int    `json:"author_id"`
	ArtistId     int    `json:"artist_id"`
	TranslatorId int    `json:"translator_id"`
	PublisherId  int    `json:"publisher_id"`
	Description  string `json:"description"`
	ViewCount    int    `json:"view_count"`
	Recommended  bool   `json:"recommended"`
}
