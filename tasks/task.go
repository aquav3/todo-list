package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task struct
type Task struct {
    ID uint64 `json:"id"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

// Gets a pcg_hash
func pcg_hash(input *uint64){
     var state uint64
     var word uint64
     state = *input * 747796405 + 2891336453
     word = ((state >> ((state >> 28) + 4)) ^ state) * 277803737
     *input = (word >> 22) ^ word
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


// Adds a task to the given task slice
func Add(tasks *[]Task) {
    var task Task
    fmt.Print("Enter the todo: ")
    reader := bufio.NewReader(os.Stdin)
    task.Title, _ = reader.ReadString('\n')
    task.ID = uint64(getID(task.Title))
    *tasks = append(*tasks, task)
}

// Lists all the tasks in the given task slice
func List(tasks *[]Task) {
    for _, task := range *tasks {
        fmt.Printf("Task: %s \n", task.Title)
        fmt.Printf("UID: %d \n", task.ID)
        fmt.Printf("Task completed: %v \n", task.Completed)
    } 
}

// Removes the task that the user provided UID matches with
func Remove(tasks *[]Task) {
    var UID uint64
    fmt.Print("Enter the UID of the task to remove: ") 
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    UID, _ = strconv.ParseUint(input, 10, 64)
        
    for i, task := range *tasks {
        if UID == task.ID {
            (*tasks)[i] = (*tasks)[len(*tasks)-1]
            (*tasks) = (*tasks)[:len(*tasks)-1]
        }
    } 
}

func WriteToFile(tasks *[]Task) {
    file, err := os.Create("tasks.json")
    if err != nil {
        fmt.Println("Failed opening the file: ", err)
        return
    } 

    data, err := json.Marshal(*tasks)
    if err != nil {
        fmt.Println("Failed encoding json: ", err)
        return
    }

    file.WriteString(string(data))
}

func ReadFromFile() []Task {
    var data []Task
    jsonData, _ := os.ReadFile("tasks.json")
    err := json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println("Failed decoding JSON: ", err)
    }

    return data
}
