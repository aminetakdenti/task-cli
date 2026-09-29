package main

import (
	"fmt"
	"strconv"
)

type App struct {
	Run func(args []string) error
}

func New() App {
	return App{
		Run: func(args []string) error {
			cli(args)
			return nil
		},
	}
}

func cli(args []string) {

	if len(args) < 2 {
		fmt.Println("Please provide at least two one for the action and one for the props.")
		return
	}

	tasksDB := Tasks{
		path: TaskPath,
	}

	command := args[0]

	if command == Add {
		data := args[1]
		newTask := tasksDB.InsertNewTask(data)
		fmt.Println()
		fmt.Println("new taks added: ", newTask)
		return
	}

	if command == Update {
		if len(args) < 3 {
			fmt.Println("Please provide at least two one for task id and one for the props.")
			return
		}
		taskID, err := strconv.Atoi(args[1])
		check(err)
		description := args[2]
		updatedTask := tasksDB.UpdateTask(taskID, description)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == Delete {
		taskId, err := strconv.Atoi(args[1])
		check(err)
		newTask := tasksDB.DeleteTask(taskId)
		fmt.Println()
		fmt.Println("Deleted Task: ", newTask)
		return
	}

	if command == MarkTodo {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusTodo)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == MarkInProgress {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusInProgress)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == MarkDone {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusDone)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == List {
		if len(args) > 1 {
			filterBy := args[1]

			if filterBy == string(StatusTodo) {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			if filterBy == string(StatusInProgress) {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			if filterBy == string(StatusDone) {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			fmt.Printf("this filter prop does not exist, for filtering use: %q %q %q", StatusTodo, StatusInProgress, StatusDone)
			return
		}
		tasks := tasksDB.ListAll()
		fmt.Println("All the tasks: ", tasks)
		return
	}

	fmt.Println("This command does not exist")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
