package main

import (
	"fmt"
)

func addTask(taskText string) {
	tasks := loadTasks()
	
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}
	
	newTask := Task{
		ID:   newID,
		Text: taskText,
	}
	
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	
	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

func removeTask(taskID int) {
	tasks := loadTasks()
	
	found := false
	var newTasks []Task
	
	for _, task := range tasks {
		if task.ID == taskID {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}
	
	if !found {
		fmt.Printf("Error: Task with ID %d not found\n", taskID)
		return
	}
	
	saveTasks(newTasks)
	fmt.Printf("Task %d removed successfully\n", taskID)
}

func listTasks() {
	tasks := loadTasks()
	
	if len(tasks) == 0 {
		fmt.Println("No tasks found. Add some tasks using 'todo add <task>'")
		return
	}
	
	fmt.Println("Your Todo List:")
	fmt.Println("---------------")
	for _, task := range tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Text)
	}
}