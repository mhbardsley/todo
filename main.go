package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mhbardsley/todo/todo"
)

// NewTodo initialises a todo file in the user's home folder
func NewTodo() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	newFilePath := filepath.Join(homeDir, ".todolist")
	_, err = os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var list todo.List
	b, err := getBytes(list)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(newFilePath, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// AddTodo adds to the existing todo list
func AddTodo() {
}

// GetTodo gets from the existing todo list
func GetTodo() {
}

// GetTodo retrieves from the existing todo list

func main() {
	command := os.Args[1]

	// process the command
	// new initialises file in home folder; add puts a new item on the todo list; get retrieves the next item on the todo list
	switch command {
	case "new":
		NewTodo()
	case "add":
		AddTodo()
	case "get":
		GetTodo()
	default:
		fmt.Println("Command not recognised")
	}
}
