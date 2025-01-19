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

## Installation

### Prerequisites

- **Go** (v1.18 or later)
- **MySQL** installed and running

### Steps to Set Up

1. **Clone the repository**:

    ```bash
    git clone https://github.com/m-golang/todo-app.git
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

4. **Set up `.env` variables for  your database connection**:
    ```bash
    MYSQL_USER_NAME=DATABASE_USER_NAME
	MYSQL_USER_PASSWORD=DATABASE_USER_PASSWORD
	MYSQL_DB_NAME=DATABASE_NAME

5. **Run the app**:

    ```bash
    go run cmd/web/main.go

