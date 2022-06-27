package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Item struct {
	task           string
	completed      bool
	creationDate   time.Time
	completionDate time.Time
	creator        string
}

type TodoList struct {
	items []Item
}

// taskExists check if a task already exists in TodoList
func (t TodoList) taskExists(task string) (bool, error) {
	for i := range t.items {
		if t.items[i].task == task {
			return true, errors.New("Task already exists it can not be added")
		}
	}
	return false, nil
}

// add appends an item to the TodoList
func (t *TodoList) add(taskName string, creator string) {

	// Check if task already exists
	_, err := t.taskExists(taskName)
	if err != nil {
		fmt.Println(err)
	} else {
		item := Item{
			task:           taskName,
			completed:      false,
			creationDate:   time.Now(),
			completionDate: time.Time{},
			creator:        creator,
		}
		t.items = append(t.items, item)
	}
}

// delete removes an element from the task list
func (t *TodoList) delete(task string) {
	var list []Item
	for i := range t.items {
		item := t.items[i]
		if item.task != task {
			t.items = append(list, item)
		}
	}
	t.items = list
}

// completeTask set a task to completed
func (t *TodoList) completeTask(task string) {
	for i := range t.items {
		if t.items[i].task == task {
			t.items[i].completed = true
			t.items[i].completionDate = time.Now()
		}
	}
}

// loadFile reads a json file containing the tasks
func (t *TodoList) loadFile(fileName string) {
	// Reading file as binary
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	// Load the binary from the file into TodoList object
	err = json.Unmarshal(file, t)
}

// storeToFile saves TodoList to json file
func (t TodoList) storeToFile(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 064)
}
