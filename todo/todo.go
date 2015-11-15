package todo

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID    int64  // Unique identifier
	Title string // Description
}

// TodoManager manages a list of todos in a sql database.
type TodoManager struct {
	db *sql.DB // Database connection
}

// Save saves the given Todo in the database.
func (t *TodoManager) Save(todo *Todo) error {
	row := t.db.QueryRow("INSERT INTO todos (title) VALUES ($1) RETURNING id", todo.Title)
	return row.Scan(&todo.ID)
}

// All returns the list of all the Tasks in the database.
func (t *TodoManager) List() ([]*Todo, error) {
	rows, err := t.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		todo := &Todo{}
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
		)
		if err != nil {
			break
		}
		todos = append(todos, todo)
	}
	return todos, err
}

// Delete deltes the Todo with the given id in the database.
func (t *TodoManager) Delete(id int64) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}

// NewTodoManager returns a TodoManager with a sql database
// setup and configured.
func NewTodoManager(driver, datasource string) (*TodoManager, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return &TodoManager{db}, nil
}

const schema = `
CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY, 
	title VARCHAR(2000)
);
`
