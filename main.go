package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var todos []Todo
var nextID = 1

func main() {
	loadTodos()
	
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		addTodo(task)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID to remove")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}
		removeTodo(id)
	case "list":
		listTodos()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID to mark as complete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}
		completeTodo(id)
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Todo List CLI")
	fmt.Println("Usage:")
	fmt.Println("  todo add <task>          - Add a new task")
	fmt.Println("  todo remove <id>         - Remove a task by ID")
	fmt.Println("  todo complete <id>       - Mark a task as complete")
	fmt.Println("  todo list                - List all tasks")
}