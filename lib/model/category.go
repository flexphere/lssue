package model

import "time"

type Category struct {
	ID        int        `db:"id" json:"id"`
	Title     string     `db:"title" json:"title"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type CategoryListParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
}

type CategoryCreateParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	Title     string `db:"title" json:"title"`
}

type CategoryDeleteParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	id        string `db:"id" json:"id"`
}
