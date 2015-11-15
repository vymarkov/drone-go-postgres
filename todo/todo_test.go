package todo

import (
	"testing"
)

var testTodo = Todo{Title: "download drone"}

var testTodos = []Todo{
	Todo{Title: "download drone"},
	Todo{Title: "setup continuous integration"},
	Todo{Title: "profit"},
}

var todos *TodoManager

func TestSave(t *testing.T) {
	setup()
	defer teardown()

	err := todos.Save(&testTodo)
	if err != nil {
		t.Errorf("Wanted to save todo, got error. %s", err)
	}
	if testTodo.ID == 0 {
		t.Errorf("Wanted todo id assignment, got 0")
	}

	after, _ := todos.List()
	if len(after) != 1 {
		t.Errorf("Wanted 1 item in the todo list, got %d todos", len(after))
	}
}

func TestList(t *testing.T) {
	setup()
	defer teardown()

	for _, todo := range testTodos {
		todos.Save(&todo)
	}

	list, err := todos.List()
	if err != nil {
		t.Errorf("Error listing todo items. %s", err)
	}
	if len(list) != len(testTodos) {
		t.Errorf("Wanted %d items in list, got %d", len(testTodos), len(list))
	}
}

func TestDelete(t *testing.T) {
	setup()
	defer teardown()

	err := todos.Save(&testTodo)
	if err != nil {
		t.Errorf("Wanted to save todo, got error. %s", err)
	}

	err = todos.Delete(testTodo.ID)
	if err != nil {
		t.Errorf("Wanted to delete todo, got error. %s", err)
	}

	after, _ := todos.List()
	if len(after) != 0 {
		t.Errorf("Wanted empty todo list, got %d todos", len(after))
	}
}

func setup() {
	todos, _ = NewTodoManager("postgres", "host=127.0.0.1 user=postgres dbname=todo sslmode=disable")
	todos.db.Exec("DELETE FROM todos")
}

func teardown() {
	todos.db.Close()
}
