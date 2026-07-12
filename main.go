package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskManager struct {
	Tasks    []Task
	Filename string
}

func (tm *TaskManager) Load() {
	data, _ := os.ReadFile(tm.Filename)
	json.Unmarshal(data, &tm.Tasks);
}

func (tm *TaskManager) Save() {
	jsonData, _ := json.MarshalIndent(tm.Tasks, "", "    ")
	os.WriteFile(tm.Filename, jsonData, 0644)
}

func (tm *TaskManager) Add(description string) {
	var newId int
	if len(tm.Tasks) == 0 {
		newId = 1
	} else {
		newId = tm.Tasks[len(tm.Tasks)-1].ID + 1
	}
	newTask := Task{
		ID : newId,
		Description : description,
		Status : "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tm.Tasks = append(tm.Tasks, newTask)
	fmt.Printf("Task added successfully (ID: %d)\n", newId)
}

func (tm *TaskManager) Update(id int, description string) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i].Description = description
			tm.Tasks[i].UpdatedAt = time.Now()
			break
		}
	}
}

func (tm *TaskManager) Delete(id int) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}

func (tm *TaskManager) ChangeStatus(id int, status string) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i].Status = status
			tm.Tasks[i].UpdatedAt = time.Now()
			break
		}
	}
}

func (tm *TaskManager) List(filterStatus string) {
	fmt.Println("ID\tStatus\t\tCreated At\t\tUpdated At\t\tDescription")
	fmt.Println("--\t------\t\t----------\t\t----------\t\t-----------")
	
	timeLayout := "02-Jan-2006 03:04 PM"

	for _, task := range tm.Tasks {
		if filterStatus == "" || task.Status == filterStatus {
			createdStr := task.CreatedAt.Format(timeLayout)
			updatedStr := task.UpdatedAt.Format(timeLayout)

			statusTab := "\t\t"
			if task.Status == "in-progress" {
				statusTab = "\t"
			}

			fmt.Printf("%d\t%s%s%s\t%s\t%s\n", 
				task.ID, 
				task.Status, 
				statusTab, 
				createdStr, 
				updatedStr, 
				task.Description,
			)
		}
	}
}

func printHelp() {
	fmt.Println("------------------------------")
	fmt.Println("Task Tracker CLI - Usage Guide")
	fmt.Println("------------------------------")
	fmt.Println("Usage:")
	fmt.Println("  task-cli <command> [arguments]\n")
	fmt.Println("Available Commands:")
	fmt.Println("  add \"<description>\"            Add a new task")
	fmt.Println("  update <id> \"<description>\"    Update a task's description")
	fmt.Println("  delete <id>                   Delete a task by its ID")
	fmt.Println("  mark-in-progress <id>         Mark a task as in-progress")
	fmt.Println("  mark-done <id>                Mark a task as done")
	fmt.Println("  list                          List all tasks")
	fmt.Println("  list todo                     List all unstarted tasks")
	fmt.Println("  list in-progress              List all ongoing tasks")
	fmt.Println("  list done                     List all completed tasks")
	fmt.Println("  help                          Show this layout guide")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Error: No command provided. Run 'task-cli help' to see the layout structure.")
		return
	}

	manager := TaskManager{
		Filename: "task.json",
	}
	manager.Load()

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task description.")
			return
		}
		manager.Add(os.Args[2])

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing task ID or new description.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		manager.Update(id, os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		manager.Delete(id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		manager.ChangeStatus(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		manager.ChangeStatus(id, "done")

	case "list":
		filter := ""
		if len(os.Args) == 3 {
			filter = os.Args[2]
		}
		manager.List(filter)

	case "help":
		printHelp()
		return

	default:
		fmt.Printf("Unknown command: '%s'. Run 'task-cli help' to see the layout structure.\n", os.Args[1])
		return
	}

	manager.Save()
}

