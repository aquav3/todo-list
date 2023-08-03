package tasks

import (
	"encoding/json"
	"errors"
	"os"
)

func (t Tasks) WriteToFile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return errors.New("Failed creating the file to save tasks to.")
    }
    data, err := json.Marshal(t.tasks) 
    if err != nil {
        return errors.New("Failed encoding the tasks.")
    }

    _, err = file.WriteString(string(data))
    if err != nil {
        return errors.New("Failed writing tasks to file.")
    }
    return nil
}

func (t Tasks) ReadFromFile(filename string) error {
    jsonData, err := os.ReadFile(filename)
    if err != nil {
        return errors.New("Failed reading tasks from file.")
    }
    err = json.Unmarshal(jsonData, &t.tasks) 
    if err != nil {
        return errors.New("Failed decoding tasks from file.") 
    }
    return nil
}
