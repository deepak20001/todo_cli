package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	// why using pointer here?
	/*
		In Go, every field of a struct must have a value. Since time.Time is a struct itself, Go initializes it with a default value (0001-01-01 00:00:00 UTC)
		That means even if a TODO is not completed, completedAt will not be nil. Instead, it will hold the default time.Time value.
		but Instead of relying on 0001-01-01 00:00:00, we can directly check for nil.
	*/
}

// slice is like vectors of c++
// This defines Todos as a new type that represents a slice ([]) of Todo structs.
// Instead of writing []Todo everywhere, you can use Todos, making the code cleaner and more readable.

type Todos []Todo

// Note:
// (*todos) dereferences the pointer to get the actual slice.
// := Declares a new variable (automatically inferring the type) & Assigns a value to it.
// Go does NOT export fields with lowercase names.

// Method to add a new TODO
func (todos *Todos) Add(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	newTodo := Todo{
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, newTodo)

	return nil
}

// Method to chek if todo index is valid
func (todos *Todos) IsIndexValid(index int) bool {
	return index >= 0 && index < len(*todos)
}

// Method to delete a todo
func (todos *Todos) Delete(index int) error {
	if !todos.IsIndexValid(index) {
		fmt.Println("Invalid index")
		return errors.New("invalid index")
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	return nil
}

// Method to update completion status
func (todos *Todos) UpdateCompletionStatus(index int) error {
	if !todos.IsIndexValid(index) {
		fmt.Println("Invalid index")
		return errors.New("invalid index")
	}

	(*todos)[index].IsCompleted = !(*todos)[index].IsCompleted
	var completedAt *time.Time
	if (*todos)[index].IsCompleted {
		now := time.Now()
		completedAt = &now
	} else {
		completedAt = nil
	}
	(*todos)[index].CompletedAt = completedAt

	return nil
}

// Method to update a todo
func (todos *Todos) Update(index int, title string) error {
	if !todos.IsIndexValid(index) {
		fmt.Println("Invalid index")
		return errors.New("invalid index")
	}

	if title == "" {
		return errors.New("title cannot be empty")
	}

	(*todos)[index].Title = title
	(*todos)[index].CompletedAt = nil

	return nil

}

// Method to print all todos
func (todos *Todos) PrintAll() error {

	t := table.New(os.Stdout)

	t.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		completedStr := "❌"
		if todo.IsCompleted {
			completedStr = "✅"
		}

		completedAtStr := "-"
		if todo.CompletedAt != nil {
			completedAtStr = todo.CompletedAt.Format(time.RFC1123)
		}

		t.AddRow(
			strconv.Itoa(i+1),
			todo.Title,
			completedStr,
			todo.CreatedAt.Format(time.RFC1123),
			completedAtStr,
		)
	}

	t.Render()

	return nil
}
