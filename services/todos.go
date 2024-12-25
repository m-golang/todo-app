package services

import (
	"database/sql"
	"strconv"
)

// SerEnv struct holds the database connection.
type SerEnv struct {
	DB *sql.DB
}

// TodoList struct represents a todo list.
type TodoList struct {
	List_id   int
	List_name string
	Tasks     []*TodoTasks
}

// TodoTasks struct represents a task within a todo list.
type TodoTasks struct {
	Id_task      int
	Task         string
	Status_task  string
	Id_list_task int
}

// GetAllLists gets all todo lists from database
func (s *SerEnv) GetAllLists() ([]*TodoList, error) {

	query := `SELECT id, todo_list_name FROM todo_lists` // Query string to get all list name

	rows, err := s.DB.Query(query) // Gets all lists from database
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lists := []*TodoList{} // New slice for all lists

	for rows.Next() { // Appends every list to lists scile if there is no error
		list := &TodoList{}

		err := rows.Scan(&list.List_id, &list.List_name)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil

}

// GetAllTasks gets all tasks from database
func (s *SerEnv) GetAllTasks() ([]*TodoTasks, error) {

	query := `SELECT id, task, is_completed, todo_list_name_id FROM tasks` // Query string to get all tasts

	rows, err := s.DB.Query(query) // Gets all tasks from database
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []*TodoTasks{} // New slice for all tasks

	for rows.Next() { // Appends every task to tasks scile if there is no error
		task := &TodoTasks{}

		err := rows.Scan(&task.Id_task, &task.Task, &task.Status_task, &task.Id_list_task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddNewList creates new tasks list to the database
func (s *SerEnv) AddNewList(listName string) error {
	query := `INSERT INTO todo_lists (todo_list_name) VALUES (?)`

	_, err := s.DB.Exec(query, listName)
	if err != nil {
		return err
	}

	return nil
}

// AddNewTask adds new task to a specific todo list in the database.
func (s *SerEnv) AddNewTask(task string, list_id int) error {
	exist := 0
	queryCheckTask := `SELECT COUNT(*) FROM todo_lists WHERE id = ?`

	err := s.DB.QueryRow(queryCheckTask, list_id).Scan(&exist)
	if err != nil {
		return err
	}

	if exist > 0 {
		query := `INSERT INTO tasks (task, todo_list_name_id) VALUES (?, ?)`

		_, err = s.DB.Exec(query, task, list_id)
		if err != nil {
			return err
		}
	} else {
		return ErrNoRecordFound
	}
	return nil
}

// ChangeTaskStatus changes status of a specific task
func (s *SerEnv) ChangeTaskStatus(taskID, isCompleted string) error {
	statusValue, err := strconv.Atoi(isCompleted)
	if err != nil {
		return ErrUnprocessableEntity
	}
	tID, err := strconv.Atoi(taskID)
	if err != nil {
		return ErrUnprocessableEntity
	}

	if statusValue == 0 {
		query := `UPDATE tasks SET is_completed = ? WHERE id = ?`

		_, err = s.DB.Exec(query, 1, tID)
		if err != nil {
			return err
		}
	} else if statusValue == 1 {
		query := `UPDATE tasks SET is_completed = ? WHERE id = ?`

		_, err = s.DB.Exec(query, 0, tID)
		if err != nil {
			return err
		}
	} else {
		return ErrUnprocessableEntity
	}

	return nil
}

// DeleteList deletes a specific task list
func (s *SerEnv) DeleteList(id int) error {
	exist := 0
	queryCheckList := `SELECT COUNT(*) FROM todo_lists WHERE id = ?`

	err := s.DB.QueryRow(queryCheckList, id).Scan(&exist)
	if err != nil {
		return err
	}

	if exist > 0 {
		_, err := s.DB.Exec("DELETE FROM todo_lists WHERE id = ?", id)
		if err != nil {
			return err
		}
	} else {
		return ErrNoRecordFound
	}

	queryCheckTask := `SELECT COUNT(*) FROM tasks WHERE todo_list_name_id = ?`
	err = s.DB.QueryRow(queryCheckTask, id).Scan(&exist)
	if err != nil {
		return err
	}

	if exist > 0 {
		_, err := s.DB.Exec("DELETE FROM tasks WHERE todo_list_name_id = ?", id)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteTask deletes a specific task in the list
func (s *SerEnv) DeleteTask(id int) error {
	exist := 0
	queryCheckTask := `SELECT COUNT(*) FROM tasks WHERE id = ?`

	err := s.DB.QueryRow(queryCheckTask, id).Scan(&exist)
	if err != nil {
		return err
	}

	if exist > 0 {
		_, err := s.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
		if err != nil {
			return err
		}
	} else {
		return ErrNoRecordFound
	}

	return nil
}
