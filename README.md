# Todo CLI Application

A simple command-line task management tool built with Go and Cobra.

## Features

- Add new tasks
- List all tasks
- Mark tasks as complete/incomplete
- Easy to use command-line interface

## Commands


## Installation

1. Clone the repository
2. Run `go install`
3. The `todo` command will be available in your terminal

## Example Usage

Add a new task

$ todo add "Write documentation" Added task 1: Write documentation
List all tasks

$ todo list [ ] 1: Write documentation
Mark task as done

$ todo done 1 Task 1 marked as done
Verify task is complete

$ todo list [✓] 1: Write documentation
Undo a completed task

$ todo undo 1 Task 1 marked as not done


## Task Storage
Tasks are stored locally in a JSON file (`tasks.json`) in your working directory.

