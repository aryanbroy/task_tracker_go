package main

import (
	"flag"
	"fmt"
	"time"
)

var fileName string = "taskData.json"

func main() {

	descPtr := flag.String("desc", "demo", "description of the task")
	statusPtr := flag.String("status", string(Todo), "describe the status of the task, choose between todo, in-progress and done ")

	flag.Parse()

	fmt.Println("Description: ", *descPtr)
	fmt.Println("Status: ", *statusPtr)

	isValidStatus := IsValidStatus(Status(*statusPtr))
	if !isValidStatus {
		fmt.Println("Invalid status, choose from todo, in-progress and done ")
		return
	}

	newTask := Task{
		Id:          1,
		Description: *descPtr,
		Status:      Status(*statusPtr),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	AppendToJson(newTask)
}
