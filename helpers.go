package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
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

func UpdateTask(taskId int, newDesc string) {
	tasks := ReadJson(fileName)

	counter := -1
	for i := range tasks {
		if tasks[i].Id == taskId {
			tasks[i].Description = newDesc
			counter = 1
		}
	}

	if counter == -1 {
		fmt.Println("No task with id: ", taskId, " exists")
		return
	}

	WriteToJson(fileName, tasks)
}

func DeleteTask(taskId int) {
	tasks := ReadJson(fileName)

	for i, _ := range tasks {
		if tasks[i].Id == taskId {
			tasks = slices.Delete(tasks, i, i+1)
			break
		}
	}

	WriteToJson(fileName, tasks)
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
