package model

import "time"

type Issue struct {
	ID        int        `db:"id" json:"id"`
	IssueID   int        `db:"issue_id" json:"issue_id"`
	Repo      string     `db:"repo" json:"repo"`
	Title     string     `db:"title" json:"title"`
	State     string     `db:"state" json:"state"`
	URL       string     `db:"url" json:"url"`
	Assignees string     `db:"assignees" json:"assignees"`
	Original  string     `db:"original" json:"original"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

type IssueListParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
}

type IssueBindParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	ID        int    `db:"id" json:"id" binding:"required"`
	TicketID  int    `db:"ticket_id" json:"ticket_id" binding:"required"`
}

type IssueUnbindParam struct {
	Boardname string `form:"board_name" json:"board_name" binding:"required"`
	ID        int    `db:"id" json:"id" binding:"required"`
	TicketID  int    `db:"ticket_id" json:"ticket_id" binding:"required"`
}
