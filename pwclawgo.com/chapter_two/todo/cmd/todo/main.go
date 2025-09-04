package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"pwclawgo.com/chapter_two/todo"
)

const todoFileName = ".todo.json"

var ErrInvalidFlag = errors.New("invalid option")

var genericErrPrefix = "An error occured:"

func main() {
	// Parsing command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be marked as completed")

	// Add some texts to the default Usage helper texts
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"%s tool. Developed By BatmanNLT as a learning experiment.\n",
			os.Args[0],
		)
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2025\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	case *list:
		// List current to do items
		fmt.Print(l)

	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		// Add the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s %w", genericErrPrefix, ErrInvalidFlag))
		os.Exit(1)
	}
}
