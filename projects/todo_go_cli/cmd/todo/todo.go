package main

import (
	_ "embed"
	"flag"
	"fmt"
	"gitlab.com/saurabh3460/devops/projects/todo_go_cli"
	"os"
)

const (
	todoFile = ".todo.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo task as complete")
	del := flag.Int("del", 0, "delete todo task")
	// task := flag.String("")
	list := flag.Bool("list", false, "list todo tasks")
	flag.Parse()

	todos := &todo.Todos{}

	if error := todos.Load(todoFile); error != nil {
		fmt.Fprintln(os.Stderr, error.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		if error := todos.Store(todoFile); error != nil {
			fmt.Fprintln(os.Stderr, error.Error())
			os.Exit(1)
		}
	case *list:
		fmt.Println(*todos)
	case *complete > 0:
		todos.Complete(*complete)
		if error := todos.Store(todoFile); error != nil {
			fmt.Fprintln(os.Stderr, error.Error())
			os.Exit(1)
		}
	case *del > 0:
		todos.Delete(*del)
		if error := todos.Store(todoFile); error != nil {
			fmt.Fprintln(os.Stderr, error.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(1)
	}

}
