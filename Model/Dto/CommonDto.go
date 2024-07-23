package Dto

type IdRequest struct {
	Id int `json:"id" binding:"required"`
}

type StringRequest struct {
	Value string `json:"value" binding:"required"`
}
