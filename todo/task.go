package todo

import (
    "fmt"
    "encoding/json"
    "os"
    "time"
)


type Task struct {
    ID          int
    Description string
    Done        bool
    CreatedAt   time.Time
    HasDeadline bool
    Deadline    time.Time
}

var tasks []Task
var lastID int


func AddTask(description string, deadline *time.Time) (Task, error) {
    // Load existing tasks
    if err := LoadTasks(); err != nil {
        return Task{}, err
    }

    // Find the highest ID
    lastID = 0
    for _, t := range tasks {
        if t.ID > lastID {
            lastID = t.ID
        }
    }

    // Create new task
    lastID++
    task := Task{
        ID:          lastID,
        Description: description,
        Done:        false,
        CreatedAt:   time.Now(),
        HasDeadline: deadline != nil,
        Deadline:    time.Time{},
    }
    if deadline != nil{
        task.Deadline = *deadline
}
    tasks = append(tasks, task)

    // Save all tasks
    if err := SaveTasks(); err != nil {
        return Task{}, err
    }

    return task, nil
}

func DeleteTask(id int) error {

    if err := LoadTasks(); err != nil {return err}

    index := -1
    for i, task := range tasks {
        if task.ID == id {
            index = i 
            break
    }

    }
    if index == -1{
        return fmt.Errorf("task with ID %d not found", id)
    }

    // Remove the task by taking everything before and after it
    tasks = append(tasks[:index], tasks[index+1:]...)
    
    return SaveTasks()


}

const taskFile = "tasks.json"

// SaveTasks saves tasks to file
func SaveTasks() error {
    data, err := json.Marshal(tasks)
    if err != nil {
        return err
    }
    return os.WriteFile(taskFile, data, 0644)
}

// LoadTasks loads tasks from file
func LoadTasks() error {
    data, err := os.ReadFile(taskFile)
    if err != nil {
        if os.IsNotExist(err) {
            return nil // It's okay if the file doesn't exist yet
        }
        return err
    }
    return json.Unmarshal(data, &tasks)
}


func GetTasks() ([]Task, error) {
    if err := LoadTasks(); err != nil {
        return nil, err
    }
    return tasks, nil
}

func MarkAsDone(id int) error {
    if err := LoadTasks(); err != nil {
        return err
    }
    
    found := false
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Done = true
            found = true
            break
        }
    }
    
    if !found {
        return fmt.Errorf("task with ID %d not found", id)
    }
    
    return SaveTasks()
}

func MarkAsUndone(id int) error {
    if err := LoadTasks(); err != nil {
        return err
    }
    
    found := false
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Done = false
            found = true
            break
        }
    }
    
    if !found {
        return fmt.Errorf("task with ID %d not found", id)
    }
    
    return SaveTasks()
}

