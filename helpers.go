package main

import (
	"errors"
	"os"
	"time"
)

func AddNewTask(description string, status Status) {

	var initTasks []Task

	if fileExists(fileName) {
		initTasks = ReadJson(fileName)
	}

	lastId := FindLastId(initTasks)

	newTask := Task{
		Id:          lastId + 1,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	AppendToJson(initTasks, newTask)

}

func FindLastId(tasks []Task) int {
	var lastId int
	for _, task := range tasks {
		lastId = task.Id
	}

	return lastId
}

func IsValidStatus(s Status) bool {
	switch s {
	case Todo, InProgress, Done:
		return true
	default:
		return false
	}
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
