package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aquav3/todo-list/tasks"
)

func printMenu() {
    fmt.Println("1: Create a task")
    fmt.Println("2: List tasks")
    fmt.Println("3: Remove a task")
    fmt.Println("4: Exit")
}

func getInput(prompt string) (string, error) {
    reader := bufio.NewReader(os.Stdin)
   
    fmt.Print(prompt)
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", errors.New("Failed reading user input.")
    }
    input = strings.TrimRight(input, "\n")

    return input, nil
}

func todo(w http.ResponseWriter, r *http.Request) {
    
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/todo", todo)
}

/*
func main() {
    var taskData tasks.Tasks
    taskData.ReadFromFile("tasks.json")
    AppLoop:
    for {
        printMenu()
        input, err := getInput("Enter option: ")
        if err != nil {
            log.Fatalln("Error: ", err)
        } 
        choice, err := strconv.Atoi(input)
        switch choice {
            case 1:
                input, err := getInput("Enter task: ")
                if err != nil {
                    log.Fatalln("Error: ", err)
                }
                taskData.Add(input)
            case 2:
                taskData.List()
            case 3:
                input, err := getInput("Enter UID: ")
                if err != nil {
                    log.Fatalln("Error: ", err)
                }
                UID, err := strconv.ParseUint(input, 10, 64)
                if err != nil {
                    log.Fatalln("Error: ", err)
                }
                taskData.Remove(UID)
            case 4:
                err := taskData.WriteToFile("tasks.json")
                if err != nil {
                    log.Fatalln("Error: ", err)
                }
                break AppLoop
        }
    }
}
*/
