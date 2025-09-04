package models

type Books struct {
	Id     int    `db:"id" json:"id" `
	Title  string `db:"title" json:"title_buku"`
	Author string `db:"author" json:"pembuat"`
}