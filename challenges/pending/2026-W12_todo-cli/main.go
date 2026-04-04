package main

import (
	"fmt"
	"os"
	"strconv"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	options := [...]string{"Todo", "In Progress", "Done"}
	if s < 0 || int(s) >= len(options) {
		return "unknown"
	}
	return options[s]
}

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	options := [...]string{"Low", "Medium", "High"}
	if p < 0 || int(p) >= len(options) {
		return "unknown"
	}
	return options[p]
}

type Task struct {
	Description string   `json:"description"`
	Priority    Priority `json:"priority"`
	Status      Status   `json:"status"`
}

type Backlog []Task

func (b *Backlog) Add(t Task) {
	*b = append(*b, t)
}

func (b Backlog) List() {
	for i, t := range b {
		fmt.Printf("%d - %s \t Priority: %s \t Status: %s", i, t.Description, t.Priority, t.Status)
	}
}
func (b *Backlog) ChangeTaskStatus(taskId int, status Status) error {
	if taskId < 0 || taskId >= len(*b) {
		return fmt.Errorf("Task with is not present in the backlog")
	}
	(*b)[taskId].Status = status
	return nil
}

func (b *Backlog) Delete(taskID int) error {
	if taskID < 0 || taskID > len(*b)-1 {
		return fmt.Errorf("Task id is not present in the backlog ")
	}
	*b = append((*b)[:taskID], (*b)[taskID+1:]...)
	return nil
}
func main() {
	fmt.Println("Todo app is running")
	var store Store
	var backlog Backlog
	err := store.Load(&backlog)
	if err != nil {
		fmt.Println("Error loading backlog", err)
		return
	}

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		fmt.Println("Command", cmd)
		switch cmd {
		case "list":
			backlog.List()
		case "add":
			if len(os.Args) < 2 {
				fmt.Print(fmt.Errorf("missing argumets for task, should at least provide description"))
				return
			} else {
				// Hate where this structure goes, nesting ifs
				t := Task{
					Description: os.Args[2],
					Priority:    Low, //default
					Status:      Todo,
				}
				if len(os.Args) > 3 {
					val, error := strconv.Atoi(os.Args[3])
					if error != nil {
						fmt.Println("Error, invalid priority")
						return
					} else {
						if val < int(Low) || val > int(High) {
							fmt.Println("Error, invalid priority")
							return
						}
						t.Priority = Priority(val)
					}
				}
				backlog.Add(t)
				fmt.Println("Backlog", backlog)
				store.Save(&backlog)
			}
		case "done":
			if len(os.Args) < 2 {
				fmt.Print(fmt.Errorf("must provide valid task id"))
				return
			}
			val := os.Args[1]
			taskId, err := strconv.Atoi(val)
			if err != nil {
				fmt.Print(fmt.Errorf("must provide valid task id"))
				return
			}
			backlog.ChangeTaskStatus(taskId, Done)
			fmt.Printf("Task %d marked Done", taskId)
			store.Save(&backlog)
		case "delete":
			if len(os.Args) < 2 {
				fmt.Print(fmt.Errorf("must provide valid task id"))
				return
			}
			val := os.Args[1]
			taskId, err := strconv.Atoi(val)

			if err != nil {
				fmt.Print(fmt.Errorf("must provide valid task id"))
			}
			err = backlog.Delete(taskId)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Printf("Task: %d successfully removed", taskId)
			store.Save(&backlog)
		}
	} else {
		fmt.Println("No arguments provided")
	}
}
