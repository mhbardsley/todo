package todo

import (
	"errors"
	"testing"
	"time"
)

var SampleTime = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

// TestValidActions runs some subtests which act on a valid sample todo list
func TestValidActions(t *testing.T) {

	// ensure that a todo item is added correctly
	t.Run("ValidAdd", func(t *testing.T) {
		addsValidItem(t, sampleList())
	})

	// ensure that a todo item is removed correctly
	t.Run("ValidRemove", func(t *testing.T) {
		removesValidItem(t, sampleList())
	})
}

// TestNoRemoveFromEmpty ensures that an appropriate error is raised when trying to remove from an empty todo list
func TestNoRemoveFromEmpty(t *testing.T) {
	emptyList := List{}

	_, err := emptyList.Get()
	if err == nil {
		t.Errorf("Expected error when popping from empty todo list")
	}
	if !errors.Is(err, ErrEmptyTodo) {
		t.Errorf("Error when popping from empty todo list was not empty")
	}
}

func addsValidItem(t *testing.T, list List) {
	item := Todo{"Emails sent back", SampleTime.Add(time.Minute)}
	list.Put(item.Goal, item.Date)

	if !Contains(list, item) {
		t.Errorf("Error when adding to todo list")
	}
}

func removesValidItem(t *testing.T, list List) {
	itemName, err := list.Get()
	expectedItem := Todo{"Bed made", SampleTime.Add(time.Second)}
	if err != nil {
		t.Errorf("Received error when removing from non-empty list")
	}
	if itemName != expectedItem.Goal {
		t.Errorf("Expected %s, got %s. Did not get the item on the todo list with the earliest deadline", expectedItem.Goal, itemName)
	} else if Contains(list, expectedItem) {
		t.Errorf("List still contains the item that should have been removed")
	}
}

func Contains(list List, t Todo) bool {
	_, ok := list[t]
	return ok
}

func sampleList() List {
	sampleList := make(List)
	sampleList[Todo{"Washing hanging up", SampleTime.Add(time.Minute)}] = nil
	sampleList[Todo{"Completed assignment", SampleTime.Add(time.Hour)}] = nil
	sampleList[Todo{"Bed made", SampleTime.Add(time.Second)}] = nil
	return sampleList
}
