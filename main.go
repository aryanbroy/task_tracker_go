package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {

	newTask := Task{
		Id:          1,
		Description: "demo 1",
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

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
