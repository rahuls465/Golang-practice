package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	addPtr := flag.String("add", "", "Add a new task")

	listPtr := flag.Bool("list", false, "List all tasks")

	completePtr := flag.Int("complete", -1, "Mark a task as complete")

	deletePtr := flag.Int("delete", -1, "Delete a task")

	flag.Parse()

	if *addPtr != "" {

		fmt.Println("Adding task:", *addPtr)
		AddTask(*addPtr)

	} else if *listPtr {

		fmt.Println("Listing all tasks")
		ListTasks()

	} else if *completePtr != -1 {

		fmt.Println("Completeing Task:", *completePtr)
		CompleteTask(*completePtr)

	} else if *deletePtr != -1 {
		fmt.Println("Delete task:")
		DeleteTask(*deletePtr)

	} else {
		fmt.Println("Invalid command")
		flag.Usage()
	}
}

const taskFile = "tasks.json"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks []Task

// Function to load tasks from a file
func loadTasks() {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		print("hit1")
		if os.IsNotExist(err) {
			tasks = []Task{}
			print("hit2")
		} else {
			log.Fatal(err)
		}
	} else {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Function to save tasks to a file
func saveTasks() {
	data, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// AddTask function to add a new task
func AddTask(title string) {
	loadTasks()
	newTask := Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	saveTasks()
	fmt.Println("Task added:", title)
}

// ListTasks function to list all tasks
func ListTasks() {
	loadTasks()
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Title, status)
	}
}

// CompleteTask function to mark a task as complete
func CompleteTask(id int) {
	loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			saveTasks()
			fmt.Println("Task completed:", tasks[i].Title)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}

// DeleteTask function to delete a task
func DeleteTask(id int) {
	loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("Task deleted:", task.Title)
			return
		}
	}
	fmt.Println("Task not found with ID:", id)
}
func print(s string){
	fmt.Println(s)
}