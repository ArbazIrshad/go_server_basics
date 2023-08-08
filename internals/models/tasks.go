package models

import (
	"database/sql"
	"time"
)

// To use this model in our handlers we need to establish a new TaskModel struct in our
// main() function and then inject it as a dependency via the application struct

// Task struct to represent
// the data for an individual task, along with a TaskModel type with methods on it to
// access and manipulate the tasks in our database.

// Define a Tsk type to hold the data for an individual TASK. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type Task struct {
	ID          int
	Title       string
	Description string
	Done        string
	Created     time.Time
	Expires     time.Time
}

// Define a TaskModel type which wraps a sql.DB connection pool.
type TaskModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *TaskModel) Insert(title string, description string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *TaskModel) Get(id int) (*Task, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *TaskModel) Latest() (*[]Task, error) {
	return nil, nil
}
