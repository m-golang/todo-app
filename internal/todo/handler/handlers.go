package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-golang/todo-app/internal/todo/forms"
	"github.com/m-golang/todo-app/internal/todo/helpers"
	services "github.com/m-golang/todo-app/internal/todo/repository"
)

// AllTodos retrieves all to-do lists and tasks, and sends them as a JSON response.
// If any error occurs while fetching the data, a server error is returned.
func AllTodos(c *gin.Context) {
	// Fetch all to-do lists along with their associated tasks
	// allLists, err := services.FetchListsWithTasks()
	// fmt.Println(allLists)
	// if err != nil {
	// 	// If an error occurs, handle it using the server error helper
	// 	helpers.ServerError(c, err)
	// 	return
	// }

	// Send the lists and their tasks in the response
	c.JSON(http.StatusOK, gin.H{
		"lists": "allLists",
	})
}

// NewList creates a new to-do list based on the data in the request body.
// If required fields are missing or an error occurs, an appropriate error is returned.
func NewList(c *gin.Context) {
	var newList forms.NewListForm

	// Bind the incoming JSON data to the NewListForm struct
	if err := c.ShouldBindJSON(&newList); err != nil {
		// If the required field 'new_list' is missing, return a missing field error
		helpers.MissingFieldErr(c, "new_list")
		return
	}

	// Add the new list using the service layer
	err := services.AddNewList(newList.Name)
	if err != nil {
		// If an error occurs, handle it using the server error helper
		helpers.ServerError(c, err)
		return
	}

	// Redirect to the main page after successfully creating the list
	c.Redirect(http.StatusSeeOther, "/")
}

// NewTask creates a new task within a specified to-do list.
// If required fields are missing or an error occurs, an appropriate error is returned.
func NewTask(c *gin.Context) {
	var newTask forms.NewTaskFrom

	// Bind the incoming JSON data to the NewTaskForm struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		// If the required fields 'new_task' and 'list_id' are missing, return a missing field error
		helpers.MissingFieldErr(c, "new_task, list_id")
		return
	}

	// Add the new task using the service layer
	err := services.AddNewTask(newTask.Task, newTask.ListID)
	if err != nil {
		// Handle errors based on specific conditions
		if errors.Is(err, helpers.ErrNoRecordFound) {
			// If no matching record is found, return a not found error
			helpers.NotFoundError(c, "List not found to add new task")
		} else {
			// For all other errors, return a server error
			helpers.ServerError(c, err)
		}
		return
	}

	// Redirect to the main page after successfully creating the task
	c.Redirect(http.StatusSeeOther, "/")
}

// TaskStatusChange updates the completion status of a task (e.g., mark as done).
// If required fields are missing or an error occurs, an appropriate error is returned.
func TaskStatusChange(c *gin.Context) {
	var statusTask forms.TaskStatusForm

	// Bind the incoming JSON data to the TaskStatusForm struct
	if err := c.ShouldBindJSON(&statusTask); err != nil {
		// If the required fields 'id_task' and 'status_task' are missing, return a missing field error
		helpers.MissingFieldErr(c, "id_task, status_task")
		return
	}

	// Change the task status using the service layer
	err := services.ChangeTaskStatus(statusTask.TaskID, statusTask.TaskStatus)
	if err != nil {
		// Handle errors based on specific conditions
		if errors.Is(err, helpers.ErrUnprocessableEntity) {
			// If the task status cannot be processed, return a bad request error
			helpers.BadRequestError(c)
		} else {
			// For all other errors, return a server error
			helpers.ServerError(c, err)
		}
		return
	}

	// Redirect to the main page after successfully updating the task status
	c.Redirect(http.StatusSeeOther, "/")
}

// DeleteTheList deletes a specific to-do list.
// If the list ID is missing or an error occurs, an appropriate error is returned.
func DeleteTheList(c *gin.Context) {
	var deleteList forms.DeleteListForm

	// Bind the incoming JSON data to the DeleteListForm struct
	if err := c.ShouldBindJSON(&deleteList); err != nil {
		// If the required field 'list_id' is missing, return a missing field error
		helpers.MissingFieldErr(c, "list_id")
		return
	}

	// Delete the to-do list using the service layer
	err := services.DeleteList(deleteList.ListID)
	if err != nil {
		// Handle errors based on specific conditions
		if errors.Is(err, helpers.ErrNoRecordFound) {
			// If no matching list is found, return a not found error
			helpers.NotFoundError(c, "List not found")
		} else {
			// For all other errors, return a server error
			helpers.ServerError(c, err)
		}
		return
	}

	// Redirect to the main page after successfully deleting the list
	c.Redirect(http.StatusSeeOther, "/")
}

// DeleteTheTask deletes a specific task.
// If the task ID is missing or an error occurs, an appropriate error is returned.
func DeleteTheTask(c *gin.Context) {
	var deleteTask forms.DeleteTaskForm

	// Bind the incoming JSON data to the DeleteTaskForm struct
	if err := c.ShouldBindJSON(&deleteTask); err != nil {
		// If the required field 'id_task' is missing, return a missing field error
		helpers.MissingFieldErr(c, "id_task")
		return
	}

	// Delete the task using the service layer
	err := services.DeleteTask(deleteTask.TaskID)
	if err != nil {
		// Handle errors based on specific conditions
		if errors.Is(err, helpers.ErrNoRecordFound) {
			// If no matching task is found, return a not found error
			helpers.NotFoundError(c, "Task not found")
		} else {
			// For all other errors, return a server error
			helpers.ServerError(c, err)
		}
		return
	}

	// Redirect to the main page after successfully deleting the task
	c.Redirect(http.StatusSeeOther, "/")
}
