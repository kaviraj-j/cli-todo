# CLI Todo App

---

### Overview:
A CLI-based todo app built using Golang.  

---

### Commands:

```bash
1. todo add <todo_title>        # Add a new todo
2. todo ls / list               # List all todos
3. todo delete <todo_id>        # Delete a todo by ID
4. todo view <todo_id>          # View details of a todo
5. todo complete <todo_id>      # Mark a todo as completed
6. todo undo <todo_id>          # Undo (mark as not completed)
7. todo help                    # Show available commands
```

---

### Installation:

**Step 1: Download the binary for your platform**

| Platform | Download |
|:---------|:---------|
| Linux | [Download `todo-linux`](https://github.com/kaviraj-j/cli-todo/releases/download/v0.1/todo-linux) |
| macOS | [Download `todo-mac`](https://github.com/kaviraj-j/cli-todo/releases/download/v0.1/todo-mac) |
| Windows | [Download `todo.exe`](https://github.com/kaviraj-j/cli-todo/releases/download/v0.1/todo.exe) |

---

**Step 2: Move the binary to your system path**

#### For Linux / macOS:

```bash
chmod +x todo-linux        # or todo-mac
sudo mv todo-linux /usr/local/bin/todo
```

#### For Windows:

- Download `todo.exe`
- Add the folder containing `todo.exe` to your **System PATH**
- Run `todo` from `cmd` or PowerShell

---

**That's it!**  
You can now run `todo` from anywhere in your terminal.

---

### Note:

- Data is stored in a local `todo-data.json` file in the working directory.
- Ensure you have write permissions in the folder where you use `todo`, or else change the data source in the code and build the binary

---
