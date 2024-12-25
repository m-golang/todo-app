# To-Do App

A simple to-do list app built with **Go** and **Gin**. You can create, manage, and delete to-do lists and tasks.

## Features

- Create and manage to-do lists
- Add, update, and delete tasks
- Mark tasks as completed or not

## Technologies

- **Go** (Backend)
- **Gin** (Web framework)
- **MySQL** (Database)
- **HTML/CSS** (Front-end)

## Installation

### Prerequisites

- **Go** (v1.18 or later)
- **MySQL** installed and running

### Steps to Set Up

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/todo-app.git
    cd todo-app

2. **Install dependencies**:

    ```bash
    go mod tidy

3. **Set up the database**:

    ```sql
    CREATE DATABASE tododb;

    USE tododb;

    CREATE TABLE todo_lists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    todo_list_name VARCHAR(255) NOT NULL
    );

    CREATE TABLE tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    is_completed TINYINT(1) DEFAULT 0,
    todo_list_name_id INT,
    FOREIGN KEY (todo_list_name_id) REFERENCES todo_lists(id)
    );

4. **Update your database connection in `cmd/web/main.go`**:

    ```bash
    dsn := flag.String("dsn", "dbsername:dbpassword@/tododb?parseTime=true", "MySQL data source name")

5. **Run the app**:

    ```bash
    go run cmd/web/main.go

### Usage

- **Home Page (GET /):** View all to-do lists and tasks.
- **Create a List (POST /list/new):** Create a new to-do list.
- **Create a Task (POST /task/new):** Add a new task to a list.
- **Change Task Status (POST /task/status):** Mark a task as done or undone.
- **Delete List (POST /list/remove):** Delete a to-do list.
- **Delete Task (POST /task/remove):** Delete a task from a list.
