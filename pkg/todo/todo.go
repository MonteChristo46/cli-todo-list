package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io/ioutil"
	"strconv"
	"time"
)

type Item struct {
	Task           string
	Completed      bool
	CreationDate   time.Time
	CompletionDate time.Time
	Creator        string
}

type TodoList struct {
	Items []Item
}

func NewTodoList(filename string) (*TodoList, error) {
	todoList := new(TodoList)
	err := todoList.loadFile(filename)
	return todoList, err
}

// taskExists check if a Task already exists in TodoList
func (t TodoList) taskExists(task string) (bool, error) {
	for i := range t.Items {
		if t.Items[i].Task == task {
			return true, errors.New("Task already exists it can not be added")
		}
	}
	return false, nil
}

// Add appends an item to the TodoList
func (t *TodoList) Add(taskName string, creator string) {

	// Check if Task already exists
	_, err := t.taskExists(taskName)
	if err != nil {
		fmt.Println(err)
	} else {
		item := Item{
			Task:           taskName,
			Completed:      false,
			CreationDate:   time.Now(),
			CompletionDate: time.Time{},
			Creator:        creator,
		}
		t.Items = append(t.Items, item)
	}
}

// Delete removes an element from the Task list
func (t *TodoList) Delete(task string) {
	var list []Item
	for i := range t.Items {
		item := t.Items[i]
		if item.Task != task {
			t.Items = append(list, item)
		}
	}
	t.Items = list
}

// CompleteTask set a Task to completed
func (t *TodoList) CompleteTask(task string) error {
	for i := range t.Items {
		if t.Items[i].Task == task {
			if t.Items[i].Completed != true {
				t.Items[i].Completed = true
				t.Items[i].CompletionDate = time.Now()
				return nil
			} else {
				return errors.New("task already set to complete")
			}
		}
	}
	return errors.New("nothing there to complete")
}

// loadFile reads a json file containing the tasks
func (t *TodoList) loadFile(fileName string) error {
	// Reading file as binary
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	// Load the binary from the file into TodoList object
	err = json.Unmarshal(file, t)
	return err
}

// StoreToFile saves TodoList to json file
func (t *TodoList) StoreToFile(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (t TodoList) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "TASK"},
			{Align: simpletable.AlignCenter, Text: "COMPLETED"},
			{Align: simpletable.AlignCenter, Text: "CREATION DATE"},
			{Align: simpletable.AlignCenter, Text: "COMPLETION DATE"},
			{Align: simpletable.AlignCenter, Text: "CREATOR"},
		},
	}

	for i := range t.Items {

		//
		var completedText string
		if t.Items[i].Completed {
			completedText = "✅"
		} else {
			completedText = "❌"
		}

		row := []*simpletable.Cell{
			{Text: strconv.Itoa(i + 1)},
			{Text: t.Items[i].Task, Align: simpletable.AlignCenter},
			{Text: completedText, Align: simpletable.AlignCenter},
			{Text: t.Items[i].CompletionDate.Format("2006-01-02 15:04")},
			{Text: t.Items[i].CreationDate.Format("2006-01-02 15:04")},
			{Text: t.Items[i].Creator},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	fmt.Println(table.String())
}
