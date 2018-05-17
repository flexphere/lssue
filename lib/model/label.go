package model

import "time"

type Label struct {
	ID        int        `db:"id" json:"id"`
	Title     string     `db:"title" json:"title"`
	Color     string     `db:"color" json:"color"`
	BGColor   string     `db:"bgcolor" json:"bgcolor"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type LabelListParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
}

type LabelCreateParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	Title     string `db:"title" json:"title"`
}

type LabelDeleteParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	id        string `db:"id" json:"id"`
}
