package main

import "time"

type Task struct {
	id          int
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)
