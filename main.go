package main

import (
	"fmt"
	"os"
	"strconv"
)

var fileName string = "taskData.json"

func main() {

	funcToDo := os.Args[1]
	fmt.Println("What to do?: ", funcToDo)

	if funcToDo == "add" {
		desc := os.Args[2]
		AddNewTask(desc, Todo)
	} else if funcToDo == "update" {
		taskId := os.Args[2]
		newDesc := os.Args[3]

		idNum, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}
		UpdateTask(idNum, newDesc)
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
