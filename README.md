## Simple Go Todo List CLI

A command-line todo list application written in Go with basic CRUD operations.

### Features
- ✅ Add new tasks
- ✅ Remove tasks by ID
- ✅ Mark tasks as complete
- ✅ List all tasks with completion status
- ✅ Persistent storage in a text file

### Installation

1. **Install Go** (if not already installed):
   ```bash
   # On macOS with Homebrew
   brew install go
   
   # On Ubuntu
   sudo apt install golang-go
   ```

2. **Build the application**:
   ```bash
   go build -o todo
   ```

### Usage

```bash
# Add a new task
./todo add "Buy groceries"
./todo add "Finish project report"

# List all tasks
./todo list

# Mark a task as complete
./todo complete 1

# Remove a task
./todo remove 2

# Show help
./todo
```

### Example Session

```bash
$ ./todo add "Learn Go programming"
Added task [1]: Learn Go programming

$ ./todo add "Write documentation"
Added task [2]: Write documentation

$ ./todo list
Todo List:
[1]   Learn Go programming
[2]   Write documentation

$ ./todo complete 1
Marked task [1] as complete

$ ./todo list
Todo List:
[1] ✓ Learn Go programming
[2]   Write documentation

$ ./todo remove 2
Removed task [2]

$ ./todo list
Todo List:
[1] ✓ Learn Go programming
```

### Data Storage

Tasks are stored in a local file `todos.txt` in the same directory. Each task is saved in the format:
```
ID|completed|task description
```

### File Structure
- `main.go` - Main entry point and command parsing
- `storage.go` - File I/O operations for persistence
- `operations.go` - Core todo list operations
- `go.mod` - Go module definition

### Requirements
- Go 1.16 or higher
- No external dependencies required