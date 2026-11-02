package main

import (
	"fmt"
	"os"
	"strconv"
)

var fileName string = "taskData.json"

func main() {

	funcToDo := os.Args[1]

	if funcToDo == "add" {
		desc := os.Args[2]
		id := AddNewTask(desc, Todo)
		fmt.Println("Task added successfully (ID: ", id, ")")
	} else if funcToDo == "update" {
		taskId := os.Args[2]
		newDesc := os.Args[3]

		idNum, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}
		UpdateTask(idNum, newDesc)
		fmt.Println("Task updated successfully")
	} else if funcToDo == "delete" {
		taskId := os.Args[2]

		idNum, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}
		DeleteTask(idNum)
		fmt.Println("Task deleted")
	} else if funcToDo == "mark-in-progress" {
		taskId := os.Args[2]

		idNum, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}

		MarkInProgress(idNum)
		fmt.Println("Done")
	} else if funcToDo == "mark-done" {
		taskId := os.Args[2]

		idNum, err := strconv.Atoi(taskId)
		if err != nil {
			fmt.Println("Error converting string to integer: ", err)
			return
		}

		MarkDone(idNum)
		fmt.Println("Done")
	} else if funcToDo == "list" {
		if len(os.Args) > 2 {
			filter := os.Args[2]
			ListFilteredTasks(Status(filter))
			return
		}
		ListTasks()
	}
}
