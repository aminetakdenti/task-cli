package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func (task *Tasks) Write(content string) {
	f, err := os.Create(task.path)
	check(err)
	defer f.Close()

	_, err = f.Write([]byte(content))
	check(err)

	fmt.Println(content)
}

func (t *Tasks) Read() []Task {
	f, err := os.Open(t.path)

	if os.IsNotExist(err) {
		return []Task{}
	}

	check(err)
	defer f.Close()
	var tasks []Task

	err = json.NewDecoder(f).Decode(&tasks)
	check(err)

	return tasks
}

func (t *Tasks) GetTasksLength() int {
	return len(t.Read())
}

func (t *Tasks) InsertNewTask(description string) Task {
	tasks := t.Read()
	newTask := Task{
		Id:          t.GetTasksLength() + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	marshalTasks, err := json.Marshal(tasks)
	check(err)

	t.Write(string(marshalTasks))

	return newTask
}

func (t *Tasks) UpdateTask(id int, description string) Task {
	tasks := t.Read()
	var task Task

	for i := range tasks {

		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			task = tasks[i]
		}
	}

	marshalTasks, err := json.Marshal(tasks)
	check(err)

	t.Write(string(marshalTasks))

	return task
}

func (t *Tasks) UpdateTaskStatus(id int, status Status) Task {
	tasks := t.Read()
	var task Task

	for i := range tasks {

		if tasks[i].Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			task = tasks[i]
		}
	}

	marshalTasks, err := json.Marshal(tasks)
	check(err)

	t.Write(string(marshalTasks))

	return task
}

func (t *Tasks) DeleteTask(id int) Task {
	tasks := t.Read()

	var deletedTask Task
	var newTasks []Task

	for _, task := range tasks {
		if task.Id == id {
			deletedTask = task
			continue
		}
		newTasks = append(newTasks, task)
	}

	marshalTasks, err := json.Marshal(newTasks)
	check(err)

	t.Write(string(marshalTasks))

	return deletedTask
}

func (t *Tasks) ListAll() []Task {
	return t.Read()
}

func (t *Tasks) ListWithFilter(status Status) []Task {
	tasks := t.Read()
	var newTasks []Task

	for i := range tasks {
		if tasks[i].Status != status {
			continue
		}
		newTasks = append(newTasks, tasks[i])
	}

	return newTasks
}
