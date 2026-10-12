package main

type Task struct {
	id          int
	description string
	status      Status
}

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)
