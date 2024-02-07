package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

const dataFile = ".todo.json"

func getDataPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, dataFile)
}

func loadTasks() []Task {
	dataPath := getDataPath()
	
	file, err := os.Open(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
		os.Exit(1)
	}

	return tasks
}

func saveTasks(tasks []Task) {
	dataPath := getDataPath()
	
	file, err := os.Create(dataPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}