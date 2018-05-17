package model

import "time"

type Ticket struct {
	ID         int        `db:"id" json:"id"`
	Title      string     `db:"title" json:"title"`
	Due        string     `db:"due" json:"due"`
	Position   string     `db:"position" json:"position"`
	Memo       string     `db:"memo" json:"memo"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at"`
	PipeID     int        `db:"pipe_id" json:"pipe_id"`
	CategoryID int        `db:"category_id" json:"category_id"`
	LabelIDS   []int      `db:"label_id" json:"label_ids"`
	IssueIDS   []int      `db:"issue_id" json:"issue_ids"`
}

type TicketListParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
}

type TicketCreateParam struct {
	Boardname  string `form:"board_name" json:"board_name" binding:"required"`
	Title      string `db:"title" json:"title" binding:"required"`
	Due        string `db:"due" json:"due"`
	Memo       string `db:"memo" json:"memo"`
	PipeID     int    `db:"pipe_id" json:"pipe_id" binding:"required"`
	CategoryID int    `db:"category_id" json:"category_id" binding:"required"`
	LabelIDS   []int  `db:"labels" json:"label_ids"`
}

type TicketUpdateParam struct {
	Boardname  string `form:"board_name" json:"board_name" binding:"required"`
	ID         int    `db:"id" json:"id" binding:"required"`
	Title      string `db:"title" json:"title" binding:"required"`
	Due        string `db:"due" json:"due"`
	Memo       string `db:"memo" json:"memo"`
	PipeID     int    `db:"pipe_id" json:"pipe_id" binding:"required"`
	CategoryID int    `db:"category_id" json:"category_id" binding:"required"`
	LabelIDS   []int  `json:"label_ids"`
}

type TicketDeleteParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	ID        int    `db:"id" json:"id" binding:"required"`
}

type TicketSortParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	PipeID    int    `json:"pipe_id" binding:"required`
	TicketIDS []int  `json:"tickets" binding:"required"`
}
