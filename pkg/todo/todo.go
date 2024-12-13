package todo

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Task struct {
	ID        int
	Name      string
	IsDone    bool
	CreatedAt string
	UpdatedAt string
}

type TodoList struct {
	Db    *sql.DB
	Tasks []Task
}

func (t *TodoList) AddTask(name string) {
	id := len(t.Tasks) + 1
	time := time.Now().String()
	task := Task{ID: id, Name: name, IsDone: false, CreatedAt: time, UpdatedAt: time}
	t.Tasks = append(t.Tasks, task)
	fmt.Printf("Task added: %s\n", name)
}

func (t *TodoList) ListTasks() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Your tasks:")
	for _, task := range t.Tasks {
		status := "Pending"
		if task.IsDone {
			status = "Done"
		}
		fmt.Printf("%d. [%s] %s %s %s\n", task.ID, status, task.Name, task.CreatedAt, task.UpdatedAt)
	}
}

func (t *TodoList) MarkAsDone(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].IsDone = true
			t.Tasks[i].UpdatedAt = time.Now().String()
			fmt.Printf("Task marked as done: %s\n", task.Name)
			return
		}
	}
	fmt.Println("Task not found.")
}

func (t *TodoList) DeleteTask(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			fmt.Printf("Task deleted: %s\n", task.Name)
			return
		}
	}
	fmt.Println("Task not found.")
}

func (t *TodoList) AddTaskDB(name string) {
	result, err := t.Db.Exec("INSERT INTO todolist (name, created_at, updated_at) VALUES (?, ?, ?)", name, time.Now(), time.Now())
	if err != nil {
		log.Fatalf("failed to insert todo: %v", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Todo added with ID: %d\n", id)
}

func (t *TodoList) ListTasksDB() {

	rows, err := t.Db.Query("SELECT * FROM todolist")
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var is_done bool
		var created_at, updated_at sql.NullTime

		if err := rows.Scan(&id, &name, &is_done, &created_at, &updated_at); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("%d. [%s] %s %s %s\n", id, name, strconv.FormatBool(is_done), created_at.Time.String(), updated_at.Time.String())
	}

}

func (t *TodoList) DeleteTaskDB(id int) {
	result, err := t.Db.Exec("DELETE FROM todolist where id = ?", id)
	if err != nil {
		log.Fatalf("failed to delete todo: %v", err)
	}
	rows, _ := result.RowsAffected()
	fmt.Printf("%d todo deleted", rows)
}

func (t *TodoList) MarkTaskDB(id int, is_done bool) {
	t.UpdateTaskDB(id, "is_done", is_done)
}

func (t *TodoList) UpdateTaskDB(id int, column string, value any) {
	result, err := t.Db.Exec(fmt.Sprintf("UPDATE todolist SET %s = ?, updated_at = ? where id = ?", column), value, time.Now(), id)
	if err != nil {
		log.Fatalf("failed to update todo: %v", err)
	}
	rows, _ := result.RowsAffected()
	fmt.Printf(" %d todo updated", rows)

}
