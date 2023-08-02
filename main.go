package main

import (
	"fmt"
	"github.com/aquav3/todo-list/tasks"
)

func printMenu() {
    fmt.Println("1: Create a task")
    fmt.Println("2: List tasks")
    fmt.Println("3: Remove a task")
    fmt.Println("4: Exit")
}

func main() {
    var choice int
    task := tasks.ReadFromFile()
    
    AppLoop:
    for {
        printMenu()
        fmt.Print("Enter option: ")
        fmt.Scanln(&choice)
        switch choice {
            case 1:
                tasks.Add(&task)
            case 2:
                tasks.List(&task)
            case 3:
                tasks.Remove(&task)
            case 4:
                tasks.WriteToFile(&task)
                break AppLoop
        }
    }
}
