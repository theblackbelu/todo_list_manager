## Simple Command-Line Todo List

A lightweight command-line todo list application written in Go.

### Installation

1. Make sure you have Go installed on your system
2. Clone or download these files to a directory
3. Navigate to the directory containing the files
4. Build the application:
   ```bash
   go build -o todo
   ```

### Usage

The application provides the following commands:

#### Add a Task
```bash
./todo add "Buy groceries"
./todo add "Finish project report"
```

#### List All Tasks
```bash
./todo list
```

#### Remove a Task
```bash
./todo remove 1
```

#### Show Help
```bash
./todo help
```

### Features

- **Persistent Storage**: Tasks are saved to a JSON file in your home directory (`~/.todo.json`)
- **Simple Interface**: Easy-to-use command-line interface
- **Task Management**: Add, remove, and list tasks with unique IDs
- **Error Handling**: Provides helpful error messages for invalid inputs

### Data Storage

Tasks are stored in a JSON file located at `~/.todo.json`. The file format is human-readable and can be manually edited if needed.

### Example Workflow

```bash
# Add some tasks
./todo add "Learn Go programming"
./todo add "Write documentation"
./todo add "Test the application"

# List tasks
./todo list
# Output:
# Your Todo List:
# ---------------
# 1. Learn Go programming
# 2. Write documentation
# 3. Test the application

# Remove a task
./todo remove 2

# List again to confirm
./todo list
# Output:
# Your Todo List:
# ---------------
# 1. Learn Go programming
# 3. Test the application
```

### Notes

- The application automatically assigns unique IDs to tasks
- Tasks are persisted between sessions
- The data file is created automatically in your home directory
- Task IDs are sequential and continue from the last assigned ID even after deletions