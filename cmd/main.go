package main

import (
	"abdqadr/todos/pkg/database"
	"abdqadr/todos/pkg/todo"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	todoList := &todo.TodoList{}
	// if err := todoList.LoadTasks(); err != nil {
	// 	fmt.Println("Error loading tasks:", err)
	// }

	db, err := database.ConnectDB()
	todoList.Db = db
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-cli <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  add <task_name>      Add a new task")
		fmt.Println("  list                 List all tasks")
		fmt.Println("  done <task_id>       Mark a task as done")
		fmt.Println("  delete <task_id>     Delete a task")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli add <task_name>")
			return
		}
		taskName := os.Args[2]
		todoList.AddTaskDB(taskName)

	case "list":
		todoList.ListTasksDB()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli done <task_id>")
			return
		}
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		todoList.MarkTaskDB(taskID, true)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli delete <task_id>")
			return
		}
		taskID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		todoList.DeleteTaskDB(taskID)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

	// if err := todoList.SaveTasks(); err != nil {
	// 	fmt.Println("Error saving tasks:", err)
	// }
}
