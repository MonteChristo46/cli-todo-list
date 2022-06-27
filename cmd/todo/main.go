package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/MonteChristo46/cli-todo-list/pkg/todo"
	"os"
)

const fileName string = "todoList.json"

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	// Defining the CLI input flags
	add := flag.String("add", "", "Adds a task to TodoList")
	complete := flag.String("complete", "", "Sets a task complete based on his name")
	deleteTask := flag.String("delete", "", "Deletes a task based on his name")
	list := flag.Bool("list", true, "Prints the TodoList in the command line")
	creator := flag.String("creator", "Unkown", "The creator that will be listed when adding task")

	// Parsing flags
	flag.Parse()

	// Initalizing TodoList
	todoList, err := todo.NewTodoList(fileName)
	if err != nil {
		fmt.Println(fileName + "not found")
	}

	// Logic for flag arguments
	if *add != "" {
		todoList.Add(*add, *creator)
		err := todoList.StoreToFile(fileName)
		handleError(err)
	}
	if *complete != "" {
		err := todoList.CompleteTask(*complete)
		handleError(err)
		err = todoList.StoreToFile(fileName)
		handleError(err)
	}

	if *deleteTask != "" {
		todoList.Delete(*deleteTask)
		err := todoList.StoreToFile(fileName)
		handleError(err)
	}
	if *list {
		todoList.Print()
	}
	if *creator != "" {
		if *add != "" {
			handleError(errors.New("creator can only be defined with valid task"))
		}
	}
}
