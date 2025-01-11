package cmd

import (
    "fmt"
    "strconv"
    "github.com/spf13/cobra"
    "todo/todo"
    "os"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [task ID]",
    Short: "Delete a task",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Error: Please provide a task ID")
            os.Exit(1)
        }
        
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Error: Invalid task ID: %s\n", args[0])
            os.Exit(1)
        }
        
        if err := todo.DeleteTask(id); err != nil {
            fmt.Printf("Error deleting task: %v\n", err)
            os.Exit(1)
        }
        
        fmt.Printf("Task %d has been deleted\n", id)
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}

