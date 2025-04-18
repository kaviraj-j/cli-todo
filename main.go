package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	COMPLETED_SYMBOL     = "✅"
	NOT_COMPLETED_SYMBOL = "⬜"
	MAX_TODO_TITLE_LEN   = 40
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("mention the operation: add, list, delete")
	}

	cmd := args[1]
	switch strings.ToLower(cmd) {
	case "help":
		fmt.Println(`Commands:
		1. todo add <todo_title>
		2. todo ls (or) todo list
		3. todo delete <todo_id>
		4. todo view <todo_id>
		5. todo complete <todo_id>
		6. todo undo <todo_id>`)

	case "add":
		if len(args) != 3 {
			log.Fatal("usage: todo add <task title>")
		}
		title := args[2]
		todos := loadTodos()

		todo := Todo{
			Title:     title,
			Completed: false,
			CreatedAt: time.Now(),
		}
		todos = append(todos, todo)
		saveTodos(todos)
		fmt.Println("✅ Added:", title)

	case "list", "ls":
		todos := loadTodos()
		if len(todos) == 0 {
			fmt.Println("📭 No todos found.")
			return
		}
		for idx, todo := range todos {
			symbol := NOT_COMPLETED_SYMBOL
			if todo.Completed {
				symbol = COMPLETED_SYMBOL
			}
			fmt.Printf("%d. %s %s - %s\n", idx+1, formatTodoTitle(todo.Title), symbol, todo.CreatedAt.Format("02 Jan 06 15:04"))
		}

	case "delete":
		todos := loadTodos()
		todo, index := findTodo(todos, args)
		todos = slices.Delete(todos, index, index+1)
		saveTodos(todos)
		fmt.Println("🗑️ Deleted:", todo.Title)

	case "complete":
		todos := loadTodos()
		todo, index := findTodo(todos, args)
		if todo.Completed {
			log.Fatal("already completed")
		}
		todo.Completed = true
		todos[index] = todo
		saveTodos(todos)
		fmt.Println("✅ Completed:", todo.Title)

	case "view":
		todos := loadTodos()
		todo, _ := findTodo(todos, args)
		symbol := NOT_COMPLETED_SYMBOL
		if todo.Completed {
			symbol = COMPLETED_SYMBOL
		}
		fmt.Printf("%s %s - %s\n", todo.Title, symbol, todo.CreatedAt.Format("02 Jan 06 15:04"))
	case "undo":
		todos := loadTodos()
		todo, index := findTodo(todos, args)
		if !todo.Completed {
			log.Fatal("todo not completed")
		}
		todo.Completed = false
		todos[index] = todo
		saveTodos(todos)
		fmt.Println("Undone:", todo.Title)
	default:
		log.Fatalf("unknown command: %s", cmd)

	}
}

func findTodo(todos []Todo, args []string) (Todo, int) {
	if len(args) != 3 {
		log.Fatal("usage: todo <option> <id>")
	}
	index, err := strconv.Atoi(args[2])

	if err != nil {
		log.Fatal("invalid ID")
	}
	index = index - 1
	if index >= len(todos) || index < 0 {
		log.Fatalf("mention the id within the range %v to %v", 1, len(todos))
	}
	return todos[index], index
}

func loadTodos() []Todo {
	data, err := os.ReadFile(getTodoDataFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}
		}
		log.Fatal("failed to read todo file:", err)
	}
	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		log.Fatal("failed to parse todo file:", err)
	}
	return todos
}

func saveTodos(todos []Todo) {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		log.Fatal("failed to encode todos:", err)
	}
	if err := os.WriteFile(getTodoDataFilePath(), data, 0644); err != nil {
		log.Fatal("failed to write todos:", err)
	}
}

func getTodoDataFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home directory: " + err.Error())
	}

	return filepath.Join(homeDir, ".todo-data.json")
}

func formatTodoTitle(todoTitle string) string {
	runes := []rune(todoTitle)
	if len(runes) > MAX_TODO_TITLE_LEN {
		return string(runes[:MAX_TODO_TITLE_LEN-3]) + "..."
	}
	return string(runes) + strings.Repeat(" ", MAX_TODO_TITLE_LEN-len(runes))
}
