package models

type Body struct {
	Id      int    `json:"id" binding:"required"`
	Message string `json:"msg"`
	Gender  string `json:"gndr"`
}