package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task to add")
			return
		}
		task := os.Args[2]
		addTask(task)
		
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task number to remove")
			return
		}
		var index int
		_, err := fmt.Sscanf(os.Args[2], "%d", &index)
		if err != nil {
			fmt.Println("Error: Please provide a valid task number")
			return
		}
		removeTask(index)
		
	case "list":
		listTasks()
		
	case "help":
		printUsage()
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Todo List Manager")
	fmt.Println("Usage:")
	fmt.Println("  todo add <task>     - Add a new task")
	fmt.Println("  todo remove <index> - Remove a task by index")
	fmt.Println("  todo list           - List all tasks")
	fmt.Println("  todo help           - Show this help message")
}