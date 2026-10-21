package main

import (
	"flag"
	"fmt"
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

	AddNewTask(*descPtr, Status(*statusPtr))
}
