package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// This Go code is handling command-line flags using the flag package.
// It allows users to pass options when running the program from the terminal, such as adding, deleting, editing, or listing todos.

// The CmdFlags struct stores command-line flag values.
type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

// Function to Parse Command-Line Flags
func NewCmdFlags() *CmdFlags {
	// Create a new CmdFlags instance
	cf := CmdFlags{}

	// Defining Command-Line Flags
	// Defines a flag -add that takes a string value (default ""). The string value will be stored in cf.Add.
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	// Defines a flag -del that takes an integer value (default -1). The integer value will be stored in cf.Del.
	flag.IntVar(&cf.Del, "del", -1, "Delete a specific todo via index")
	// Defines a flag -edit that takes a string value (default ""). The string value will be stored in cf.Edit.
	flag.StringVar(&cf.Edit, "edit", "", "Edit a specific todo via index & it's new title in format index:newTitle")
	// Defines a flag -toggle that takes an integer value (default 0). The integer value will be stored in cf.Toggle.
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a specific todo via index")
	// Defines a flag -list that takes a boolean value (default false). The boolean value will be stored in cf.List.
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	// This parses the command-line flags and stores their values in cf.
	flag.Parse()

	// The function returns a pointer to CmdFlags with the parsed values.
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.PrintAll()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	case cf.Edit != "":
		// in format index: newTitle
		parts := strings.Split(cf.Edit, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid format. Use index:newTitle")
			return
		}
		index, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Invalid index")
			return
		}
		todos.Update(index, strings.TrimSpace(parts[1]))
	case cf.Toggle != -1:
		todos.UpdateCompletionStatus(cf.Toggle)

	default:
		fmt.Println("Invalid command")
	}
}
