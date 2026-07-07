# Go Task Tracker CLI

A lightweight, **zero-dependency** command-line task manager built from scratch in **Go** for managing and tracking daily tasks. This project follows the core guidelines specified on **roadmap.sh**.

The application allows users to **add, update, delete, and track the status of tasks** directly from the terminal. All task data is stored locally in a structured **JSON** file with automatic creation and update timestamps.

---

## Features

* ➕ Add new tasks
* ✏️ Update existing task descriptions
* 🗑️ Delete tasks
* 🚀 Mark tasks as **In Progress**
* ✅ Mark tasks as **Done**
* 📋 List all tasks
* 🔍 Filter tasks by status
* 💾 Persistent local JSON storage
* 🕒 Automatic creation & update timestamps
* ⚡ Built using only Go's standard library (No external dependencies)

---

# Architecture Overview

The project is designed using standard object-oriented principles in Go while relying solely on the Go standard library.

### Data Modeling (`Task`)

A custom Go struct representing each task with the following properties:

* Unique incremental ID
* Description
* Status
* `time.Time` based creation timestamp
* `time.Time` based update timestamp

These fields are automatically serialized into JSON.

---

### State Management (`TaskManager`)

A dedicated service abstraction responsible for:

* Managing the in-memory task slice
* Reading and writing `task.json`
* JSON serialization/deserialization
* Handling all CRUD operations

Implemented using:

* `os`
* `encoding/json`

---

### Command Routing (`main`)

The CLI parser performs:

* Argument validation
* Command lookup
* Handler execution
* Runtime safety checks

This approach keeps the application lightweight while avoiding external CLI frameworks.

---

# Prerequisites

To build or run this project, install:

* **Go 1.16 or higher**

Download Go from:

https://go.dev/dl/

---

# Local Setup & Installation

## 1️⃣ Clone the Repository

```bash
git clone https://github.com/YOUR_GITHUB_USERNAME/go-task-tracker-cli.git

cd go-task-tracker-cli
```

---

## 2️⃣ Build the Executable

### Windows

```bash
go build -o task-cli.exe main.go
```

### Linux / macOS

```bash
go build -o task-cli main.go

chmod +x task-cli
```

---

## 3️⃣ Global Installation (Optional)

This allows the command to be executed from any directory.

### Windows

Add your project directory to your **Path Environment Variable**, then restart your terminal.

### Linux / macOS

```bash
sudo mv task-cli /usr/local/bin/
```

---

# Usage Guide

## Show Help

Displays all supported commands.

```bash
task-cli help
```

---

## Add a New Task

Creates a new task with the default status **todo**.

```bash
task-cli add "Buy groceries and prep meals"
```

Output

```text
Task added successfully (ID: 1)
```

---

## Update a Task Description

Updates the description of an existing task.

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

---

## Delete a Task

Removes a task permanently.

```bash
task-cli delete 1
```

---

## Update Task Status

Move tasks through their workflow using dedicated commands.

Mark as In Progress

```bash
task-cli mark-in-progress 1
```

Mark as Done

```bash
task-cli mark-done 1
```

---

## List Tasks

### List All Tasks

```bash
task-cli list
```

### List Only Todo Tasks

```bash
task-cli list todo
```

### List Only In Progress Tasks

```bash
task-cli list in-progress
```

### List Only Completed Tasks

```bash
task-cli list done
```

---

# Storage & Persistence

Task data is stored locally in a file named **`task.json`**, which is automatically created on the first execution.

Example:

```json
[
  {
    "id": 1,
    "description": "Buy groceries and cook dinner",
    "status": "in-progress",
    "created_at": "2026-07-08T01:05:22.1234567+05:30",
    "updated_at": "2026-07-08T01:08:45.9876543+05:30"
  }
]
```

---

# Project Structure

```text
.
├── main.go
├── task.json
├── README.md
```

---

# uilt With

* Go
* encoding/json
* os
* time

No third-party libraries or CLI frameworks are used.

---

## Project Specification

This project is built by following the **Task Tracker** project specification from **roadmap.sh**.

 **Project Page:** https://roadmap.sh/projects/task-tracker

---

# License

This project is intended for learning purposes and follows the specifications provided by roadmap.sh.
