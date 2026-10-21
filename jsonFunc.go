package main

import (
	"encoding/json"
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
	return tasks
}

func AppendToJson(initTasks []Task, newTask Task) {

	newList := append(initTasks, newTask)

	// dataBytes, err := json.Marshal(newList)
	// if err != nil {
	// 	log.Fatalln("Error marshaling data: ", err)
	// }

	WriteToJson(fileName, newList)
}
