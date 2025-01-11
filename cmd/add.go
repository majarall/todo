package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "todo/todo"
    "os"
    "time"
)

var deadlineFlag string

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

        var deadline *time.Time
        if deadlineFlag != "" {
            parsedTime, err := time.Parse("2006-01-02", deadlineFlag)
            if err != nil {
                fmt.Printf("Error: Invalid deadline format. Use YYYY-MM-DD\n")
                os.Exit(1)
            }
            deadline = &parsedTime
        }

        task, err := todo.AddTask(args[0], deadline)
        if err != nil {
            fmt.Printf("Error adding task: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("Added task %d: %s\n", task.ID, task.Description)
}

func init() {
    rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringVar(&deadlineFlag, "deadline", "", "Set deadline (format: YYYY-MM-DD)")
}

