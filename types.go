package main

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)
