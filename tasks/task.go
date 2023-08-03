package tasks

import (
	"fmt"
)

// Task struct
type task struct {
    ID uint64 `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

type Tasks struct {
    tasks []task
}

// Gets the unique ID for a task based on the task content
func getID(content string) uint64 {
    var ID uint64
    ID = 0
    for i, char := range content {
        ID += uint64(char+1) * uint64(i)    
    }
    pcg_hash(&ID)
    return ID
}


// Adds a task with the content provided
func (t *Tasks) Add(content string) {
    tsk := task {
        Title: content,
        ID: uint64(getID(content)),
        Completed: false,
    }
    t.tasks = append(t.tasks, tsk)
}

// Lists all the tasks in the given task slice
func (t *Tasks) List() {
    for i, task := range t.tasks { 
        fmt.Println(i)
        fmt.Println("Task: ", task.Title)
        fmt.Println("UID: ", task.ID)
        fmt.Println("Task completed: ", task.Completed)
    }   
}

// Removes the task that the user provided UID matches with
func (t *Tasks) Remove(UID uint64) {
    for i, task := range t.tasks {
        if UID == task.ID {
            t.tasks[i] = t.tasks[len(t.tasks)-1]
            t.tasks = t.tasks[:len(t.tasks)-1]
        }
    }    
}

