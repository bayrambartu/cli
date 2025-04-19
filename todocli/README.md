# Todo CLI with Cobra

This is a simple command-line Todo List application built in Go using the [Cobra](https://github.com/spf13/cobra) framework.

It demonstrates how to build modular CLI applications with commands like `add`, `list`, and `delete`, while saving data to a JSON file.

---

## Features

- Add todo items
- List all todos
- Delete todos by index
- Stores todos in a local `todos.json` file
- Easy to extend with additional commands (e.g., complete, update, clear)

---

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/bayrambartu/todo-cli.git
cd todo-cli
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Run the app

```bash
go run main.go
```

---

## Commands

### 1. Add a todo

```bash
go run main.go add "Learn Go"
```

### 2. List all todos

```bash
go run main.go list
```

### 3. Delete a todo by index

```bash
go run main.go delete 1
```

---

## Project Structure

```
todo-cli/
├── cmd/
│   ├── add.go
│   ├── list.go
│   ├── delete.go
│   ├── utils.go
│   └── root.go
├── todos.json
└── main.go
```