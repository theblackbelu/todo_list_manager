package main

import "fmt"

func addTodo(task string) {
	todo := Todo{
		ID:   nextID,
		Task: task,
		Done: false,
	}
	todos = append(todos, todo)
	nextID++
	saveTodos()
	fmt.Printf("Added task [%d]: %s\n", todo.ID, task)
}

func removeTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			fmt.Printf("Removed task [%d]\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func completeTodo(id int) {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Done = true
			saveTodos()
			fmt.Printf("Marked task [%d] as complete\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found\n", id)
}

func listTodos() {
	if len(todos) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("Todo List:")
	for _, todo := range todos {
		status := " "
		if todo.Done {
			status = "✓"
		}
		fmt.Printf("[%d] %s %s\n", todo.ID, status, todo.Task)
	}
}