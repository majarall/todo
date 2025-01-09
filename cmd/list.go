package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "todo/todo"
    "os"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks, err := todo.GetTasks()
        if err != nil {
            fmt.Printf("Error getting tasks: %v\n", err)
            os.Exit(1)
        }

        if len(tasks) == 0 {
            fmt.Println("No tasks to display")
            return
        }

        for _, task := range tasks {
            status := " "
            if task.Done {
                status = "✓"
            }
            fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Description)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}

