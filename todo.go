package todo

import (
	"errors"
	"time"
)

// ErrEmptyTodo should be thrown when trying to remove from an empty list
var ErrEmptyTodo = errors.New("error when popping from empty todo list was not empty")

// Todo struct provides the basic structure of a Todo
type Todo struct {
	Goal string
	Date time.Time
}

// List provides a set of Todos
type List map[Todo]interface{}

// Get pops from the todo list the earliest goal, returning an error if the list is empty
func (list List) Get() (s string, err error) {
	var bestTodo *Todo
	for key := range list {
		if bestTodo == nil || key.Date.Before(bestTodo.Date) {
			bestTodo = &key
		}
	}
	if bestTodo == nil {
		return s, ErrEmptyTodo
	}
	delete(list, *bestTodo)
	return bestTodo.Goal, nil
}

// Put adds to the list
func (list List) Put(s string, t time.Time) {
	list[Todo{s, t}] = nil
}
