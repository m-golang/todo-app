package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/m-golang/todo-app/internal/todo/helpers"
)

var DB *sql.DB

// TodoList struct represents a todo list.
type TodoList struct {
	List_id   int          // Unique identifier for the list
	List_name string       // Name of the todo list
	Tasks     []*TodoTasks // Tasks associated with the todo list
}

// TodoTasks struct represents a task within a todo list.
type TodoTasks struct {
	Id_task      int    // Unique identifier for the task
	Task         string // Description of the task
	Status_task  string // Task completion status
	Id_list_task int    // Foreign key to identify the list the task belongs to
}

// RetrieveAllLists fetches all todo lists from the database.
// Returns a slice of TodoList pointers and an error (if any).
func RetrieveAllLists() ([]*TodoList, error) {

	query := `SELECT id, todo_list_name FROM todo_lists` // Query string to get all list name

	rows, err := DB.Query(query) // Gets all lists from database
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := []*TodoList{} // New slice for all lists

	for rows.Next() { // Appends every list to lists scile if there is no error
		list := &TodoList{}

		// Scan row values into list fields
		err := rows.Scan(&list.List_id, &list.List_name)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list) // Append the list to the slice
	}

	if err = rows.Err(); err != nil {
		return nil, err // Return error if there is an issue after scanning rows
	}

	return lists, nil // Return the slice of lists

}

// RetrieveAllTasks fetches all tasks from the database.
// Returns a slice of TodoTasks pointers and an error (if any).
func RetrieveAllTasks() ([]*TodoTasks, error) {

	query := `SELECT id, task, is_completed, todo_list_name_id FROM tasks` // Query string to get all tasts

	rows, err := DB.Query(query) // Gets all tasks from database
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []*TodoTasks{} // New slice for all tasks

	for rows.Next() { // Iterate through each row (task)
		task := &TodoTasks{}

		// Scan row values into task fields
		err := rows.Scan(&task.Id_task, &task.Task, &task.Status_task, &task.Id_list_task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task) // Append the task to the slice
	}

	if err = rows.Err(); err != nil {
		return nil, err // Return error if there is an issue after scanning rows
	}

	return tasks, nil // Return the slice of tasks
}

// AddNewList creates a new todo list in the database.
// Takes the list name as input and returns an error (if any).
func AddNewList(listName string) error {
	query := `INSERT INTO todo_lists (todo_list_name) VALUES (?)` // Query to insert a new todo list

	_, err := DB.Exec(query, listName) // Execute the query
	if err != nil {
		return fmt.Errorf("failed to insert new todo list: %w", err)
	}

	return nil
}

// AddNewTask adds a new task to a specific todo list in the database.
// Takes task description and list ID as input, returns an error (if any).
func AddNewTask(task string, list_id int) error {
	exist := 0
	queryCheckTask := `SELECT COUNT(*) FROM todo_lists WHERE id = ?` // Query to check if the list exists

	err := DB.QueryRow(queryCheckTask, list_id).Scan(&exist) // Check if list exists by ID
	if err != nil {
		return fmt.Errorf("failed to count todo list: %w", err) // Return error if check fails
	}

	if exist > 0 { // If list exists, proceed with task insertion
		query := `INSERT INTO tasks (task, todo_list_name_id) VALUES (?, ?)` // Query to insert task

		_, err = DB.Exec(query, task, list_id) // Execute insertion query
		if err != nil {
			return fmt.Errorf("failed to insert new todo task: %w", err) // Return error if insertion fails
		}
	} else {
		return helpers.ErrNoRecordFound // Return custom error if list is not found
	}
	return nil // Return nil if task was added successfully
}

// ChangeTaskStatus updates the status of a specific task (mark it as completed or not).
// Takes task ID and status (0 for completed, 1 for not completed) as input, returns an error (if any).
func ChangeTaskStatus(taskID int, isCompleted string) error {
	statusValue, err := strconv.Atoi(isCompleted) // Convert status to integer
	if err != nil {
		return helpers.ErrUnprocessableEntity // Return error if conversion fails
	}

	// If status is 0 (completed), update the task status to 1 (completed)
	if statusValue == 0 {
		query := `UPDATE tasks SET is_completed = ? WHERE id = ?` // Query to update task status

		_, err = DB.Exec(query, 1, taskID) // Execute status update
		if err != nil {
			return fmt.Errorf("failed to set todo task status: %w", err) // Return error if update fails
		}
	} else if statusValue == 1 {
		query := `UPDATE tasks SET is_completed = ? WHERE id = ?` // Query to update task status

		_, err = DB.Exec(query, 0, taskID) // Execute status update
		if err != nil {
			return fmt.Errorf("failed to set todo task status: %w", err) // Return error if update fails
		}
	} else {
		return helpers.ErrUnprocessableEntity // Return error if status is invalid
	}

	return nil // Return nil if task status was successfully changed
}

// DeleteList deletes a specific todo list and its associated tasks from the database.
// Takes list ID as input and returns an error (if any).
func DeleteList(id int) error {
	// Check if list exists
	exist := 0
	queryCheckList := `SELECT COUNT(*) FROM todo_lists WHERE id = ?`

	err := DB.QueryRow(queryCheckList, id).Scan(&exist) // Check if list exists
	if err != nil {
		return fmt.Errorf("failed to count todo list: %w", err) // Return error if query fails
	}

	if exist > 0 { // If list exists, proceed with deletion
		_, err := DB.Exec("DELETE FROM todo_lists WHERE id = ?", id) // Query to delete list
		if err != nil {
			return fmt.Errorf("failed to delete todo list: %w", err) // Return error if deletion fails
		}
	} else {
		return helpers.ErrNoRecordFound // Return error if list is not found
	}

	// Check and delete associated tasks
	exist = 0
	queryCheckTask := `SELECT COUNT(*) FROM tasks WHERE todo_list_name_id = ?`
	err = DB.QueryRow(queryCheckTask, id).Scan(&exist) // Check if tasks exist for the list
	if err != nil {
		return fmt.Errorf("failed to count todo task: %w", err) // Return error if query fails
	}

	if exist > 0 { // If tasks exist, delete them
		_, err := DB.Exec("DELETE FROM tasks WHERE todo_list_name_id = ?", id) // Query to delete tasks
		if err != nil {
			return fmt.Errorf("failed to delete todo task: %w", err) // Return error if deletion fails
		}
	}

	return nil // Return nil if list and tasks were successfully deleted
}

// DeleteTask deletes a specific task from the database.
// Takes task ID as input and returns an error (if any).
func DeleteTask(id int) error {
	// Check if task exists
	exist := 0
	queryCheckTask := `SELECT COUNT(*) FROM tasks WHERE id = ?`

	err := DB.QueryRow(queryCheckTask, id).Scan(&exist) // Check if task exists
	if err != nil {
		return fmt.Errorf("failed to count todo task: %w", err) // Return error if query fails
	}

	if exist > 0 { // If task exists, proceed with deletion
		_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id) // Query to delete task
		if err != nil {
			return fmt.Errorf("failed to delete todo task: %w", err) // Return error if deletion fails
		}
	} else {
		return helpers.ErrNoRecordFound // Return error if task is not found
	}

	return nil // Return nil if task was successfully deleted
}
