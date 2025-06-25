package models

import (
	"database/sql"
	"fmt"

	"github.com/Yashwanth1906/Go-Todo/backend/pkg/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Status int

const (
	Pending Status = iota
	Completed
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Completed:
		return "completed"
	default:
		return "pending"
	}
}

func StatusFromString(s string) Status {
	switch s {
	case "completed":
		return Completed
	case "pending":
		return Pending
	default:
		return Pending
	}
}

func init() {
	config.Connect()
	db = config.GetDB()
	// dropTable() // whenever i ever migrate i need to drop the table if i want to....
	migrateDB()
}

func dropTable() {
	query := `DROP TABLE IF EXISTS tasks`
	if _, err := db.Exec(query); err != nil {
		panic("Failed to drop table: " + err.Error())
	}
	enumDropQuery := `DROP TYPE IF EXISTS task_status`
	if _, err := db.Exec(enumDropQuery); err != nil {
		panic("Failed to drop enum type: " + err.Error())
	}
	fmt.Println("Dropped successfully...")
}

func migrateDB() {
	enumQuery := `DO $$ BEGIN
		CREATE TYPE task_status AS ENUM ('pending', 'completed');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`

	tableQuery := `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description VARCHAR(300),
		status task_status DEFAULT 'pending'
	)`
	
	if _, err := db.Exec(enumQuery); err != nil {
		panic("Failed to create enum type: " + err.Error())
	}
	if _, err := db.Exec(tableQuery); err != nil {
		panic("Failed to create table: " + err.Error())
	}
	fmt.Println("Migrated Succesfully")
}

func AddTask(task *Task) (*Task, error) {
	query := `INSERT INTO tasks (name, description, status) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, task.Name, task.Description, task.Status.String()).Scan(&task.ID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetTasks() []Task {
	query := `SELECT id, name, description, status FROM tasks ORDER BY id`
	rows, err := db.Query(query)
	if err != nil {
		return []Task{}
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
		var statusStr string
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &statusStr)
		if err != nil {
			continue
		}
		task.Status = StatusFromString(statusStr)
		tasks = append(tasks, task)
	}
	return tasks
}

func DeleteTask(id int64) (Task, error) {
	var task Task
	var statusStr string
	selectQuery := `SELECT id, name, description, status FROM tasks WHERE id = $1`
	err := db.QueryRow(selectQuery, id).Scan(&task.ID, &task.Name, &task.Description, &statusStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, fmt.Errorf("task not found")
		}
		return Task{}, err
	}

	task.Status = StatusFromString(statusStr)
	deleteQuery := `DELETE FROM tasks WHERE id = $1`
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		return Task{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Task{}, err
	}
	if rowsAffected == 0 {
		return Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}

func GetTaskById(id int64) (Task, error) {
	var task Task
	var statusStr string
	query := `SELECT id, name, description, status FROM tasks WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Name, &task.Description, &statusStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, fmt.Errorf("task not found")
		}
		return Task{}, err
	}
	task.Status = StatusFromString(statusStr)
	return task, nil
}

func UpdateStatusById(id int64) (Task, error) {
	var task Task
	var statusStr string
	selectQuery := `SELECT id, name, description, status FROM tasks WHERE id = $1`
	err := db.QueryRow(selectQuery, id).Scan(&task.ID, &task.Name, &task.Description, &statusStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return Task{}, fmt.Errorf("task not found")
		}
		return Task{}, err
	}

	task.Status = StatusFromString(statusStr)
	var newStatus Status
	if task.Status == Completed {
		newStatus = Pending
	} else {
		newStatus = Completed
	}
	updateQuery := `UPDATE tasks SET status = $1 WHERE id = $2`
	result, err := db.Exec(updateQuery, newStatus.String(), id)
	if err != nil {
		return Task{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Task{}, err
	}
	if rowsAffected == 0 {
		return Task{}, fmt.Errorf("task not found")
	}
	task.Status = newStatus
	return task, nil
}
