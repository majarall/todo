package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "todo/todo"
    "os"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new task",
    Run:addRun,
}

func addRun(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Error: Please provide a task description")
            os.Exit(1)
        }
        description := args[0]
        task, err := todo.AddTask(description)
        if err != nil {
            fmt.Printf("Error adding task: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
}

func init() {
    rootCmd.AddCommand(addCmd)
}

