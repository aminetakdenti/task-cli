package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	CommandAdd    = "add"
	CommandUpdate = "update"
	CommandDelete = "delete"
	CommandList   = "list"

	CommandMarkTodo       = "mark-todo"
	CommandMarkInProgress = "mark-in-progress"
	CommandMarkDone       = "mark-done"
)

const (
	FilterTodo       = "todo"
	FilterInProgress = "in-progress"
	FilterDone       = "done"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide at least two one for the action and one for the props.")
		return
	}

	tasksDB := Tasks{
		path: TaskPath,
	}

	args := os.Args[1:len(os.Args)]
	command := args[0]

	if command == CommandAdd {
		data := args[1]
		newTask := tasksDB.InsertNewTask(data)
		fmt.Println()
		fmt.Println("new taks added: ", newTask)
		return
	}

	if command == CommandUpdate {
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

	if command == CommandDelete {
		taskId, err := strconv.Atoi(args[1])
		check(err)
		newTask := tasksDB.DeleteTask(taskId)
		fmt.Println()
		fmt.Println("Deleted Task: ", newTask)
		return
	}

	if command == CommandMarkTodo {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusTodo)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == CommandMarkInProgress {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusInProgress)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == CommandMarkDone {
		taskID, err := strconv.Atoi(args[1])
		check(err)
		updatedTask := tasksDB.UpdateTaskStatus(taskID, StatusDone)
		fmt.Println()
		fmt.Println("The updated Task: ", updatedTask)
		return
	}

	if command == CommandList {
		if len(args) > 1 {
			filterBy := args[1]

			if filterBy == FilterTodo {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			if filterBy == FilterInProgress {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			if filterBy == FilterDone {
				fmt.Printf("All %q the tasks: %v\n", filterBy, tasksDB.ListWithFilter(Status(filterBy)))
				return
			}

			fmt.Printf("this filter prop does not exist, for filtering use: %q %q %q", FilterTodo, FilterInProgress, FilterDone)
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
