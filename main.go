package main

import "fmt"

func main() {
	todosList := Todos{}
	storage := NewStorage[Todos]("todos.json")
	err := storage.Load(&todosList)
	if err != nil {
		fmt.Println("Error loading todos:", err)
	}

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todosList)
	storage.Save(todosList)

}
