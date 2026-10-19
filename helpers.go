package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func WriteToJson(fileName string, newTask []Task) {
	file, err := os.Create("taskData.json")
	if err != nil {
		log.Fatalln("Error writing into json file: ", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(newTask)
	if err != nil {
		log.Fatalln("Error encoding newTask: ", err)
	}

}

func ReadJson(fileName string) []Task {
	bytData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}

	tasks := []Task{}

	json.Unmarshal(bytData, &tasks)
	fmt.Println("Reading initial data: ", tasks)
	return tasks
}

func AppendToJson(task Task) {

	var initTasks []Task

	if fileExists(fileName) {
		initTasks = ReadJson(fileName)
	}

	newList := append(initTasks, task)

	// dataBytes, err := json.Marshal(newList)
	// if err != nil {
	// 	log.Fatalln("Error marshaling data: ", err)
	// }

	WriteToJson(fileName, newList)
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
