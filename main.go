package main

import (
	"fmt"
	"os"
)

var fileName string = "taskData.json"

func main() {

	funcToDo := os.Args[1]
	fmt.Println("What to do?: ", funcToDo)

	var desc string
	if funcToDo == "add" {
		desc = os.Args[2]
		AddNewTask(desc, Todo)
	}

	// descPtr := flag.String("desc", "demo", "description of the task")
	// statusPtr := flag.String("status", string(Todo), "describe the status of the task, choose between todo, in-progress and done ")
	//
	// flag.Parse()
	//
	// fmt.Println("Description: ", *descPtr)
	// fmt.Println("Status: ", *statusPtr)
	//
	// isValidStatus := IsValidStatus(Status(*statusPtr))
	// if !isValidStatus {
	// 	fmt.Println("Invalid status, choose from todo, in-progress and done ")
	// 	return
	// }
	//
	// AddNewTask(*descPtr, Status(*statusPtr))
}
