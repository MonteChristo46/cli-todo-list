package main

import (
	"flag"
)

const fileName string = "todoList.json"

func main() {

	// Defining the CLI input flags
	add := flag.String("task", "", "Adds a task to TodoList")
	complete := flag.String("complete", "", "Sets a task complete based on his name")
	delete := flag.String("delete", "", "Deletes a task based on his name")

	// Initalizing TodoList
	todoList := TodoList{}

	switch {
	case *add != "":
		todoList.add(add, "Daniel")
		return
	}
}
