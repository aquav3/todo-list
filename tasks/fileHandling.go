package tasks

import (
    "os"
    "encoding/json"
    "fmt"
)

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
