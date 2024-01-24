package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dataFile = "todos.txt"

func loadTodos() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return // File doesn't exist yet, that's fine
		}
		fmt.Printf("Error loading todos: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) >= 3 {
			id, _ := strconv.Atoi(parts[0])
			done, _ := strconv.ParseBool(parts[1])
			task := parts[2]
			
			todo := Todo{
				ID:   id,
				Task: task,
				Done: done,
			}
			todos = append(todos, todo)
			
			if id >= nextID {
				nextID = id + 1
			}
		}
	}
}

func saveTodos() {
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		return
	}
	defer file.Close()

	for _, todo := range todos {
		line := fmt.Sprintf("%d|%t|%s\n", todo.ID, todo.Done, todo.Task)
		file.WriteString(line)
	}
}