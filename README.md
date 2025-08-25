
# Todo CLI in Go

A simple command-line todo manager written in Go.
Supports adding, editing, deleting, toggling, and listing tasks, saved to a JSON file.

## Features
- Add, list, edit, delete, toggle tasks
- Stores data locally in `todos.json`
- Pretty CLI output with table view

## Installation
Clone the repo and build the binary:
```bash
git clone https://github.com/<yourusername>/todo_app.git
cd todo_app
go build -o todo ./cmd/todo
```
Optionally, move the binary to a folder in your PATH (e.g., /usr/local/bin) to use it globally:

## Usage
Once built, you can use the CLI like this:
# Add a new task
todo add "Buy groceries"
# List all tasks
todo list
# Toggle a task's completion status
todo toggle 0
# Edit a task's title
todo edit 0 "Buy groceries and cook dinner"
# Delete a task
todo delete 0

# Print tasks in a table format
todo print
