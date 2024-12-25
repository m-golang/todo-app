package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-golang/todo-app/services"
)

// AllTodos retrieves all to-do lists and tasks and sends them in the response.
func (appEnv *appEnv) AllTodos(c *gin.Context) {
	// Fetch all to-do lists from the service layer
	allLists, err := appEnv.services.GetAllLists()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	// Fetch all tasks from the service layer
	allTasks, err := appEnv.services.GetAllTasks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	// Create a map to group tasks by their associated to-do list ID
	tasksByList := make(map[int][]*services.TodoTasks)

	// Iterate over all tasks and group them by the list they belong to
	for _, task := range allTasks {
		tasksByList[task.Id_list_task] = append(tasksByList[task.Id_list_task], task)
	}

	// Iterate over all to-do lists and assign the corresponding tasks from the 'tasksByList' map
	for i := range allLists {
		listID := allLists[i].List_id
		allLists[i].Tasks = tasksByList[listID]
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"lists": allLists, // Pass the lists (with their associated tasks) to the template
	})
}

// NewList creates a new to-do list based on the data in the request body.
func (appEnv *appEnv) NewList(c *gin.Context) {
	var newList NewListForm

	if err := c.ShouldBind(&newList); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Field cannot be blank",
		})
		return
	}

	err := appEnv.services.AddNewList(newList.Name)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

// NewTask creates a new task in a specified to-do list.
func (appEnv *appEnv) NewTask(c *gin.Context) {
	var newTask NewTaskFrom

	if err := c.ShouldBind(&newTask); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Field cannot be blank",
		})
		return
	}

	err := appEnv.services.AddNewTask(newTask.Task, newTask.ListID)
	if err != nil {
		if errors.Is(err, services.ErrNoRecordFound) {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": http.StatusText(http.StatusBadRequest),
			})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.Redirect(http.StatusSeeOther, "/")

}

// TaskStatusChange updates the completion status of a task (e.g., mark as done).
func (appEnv *appEnv) TaskStatusChange(c *gin.Context) {
	var statusTask TaskStatusForm

	if err := c.ShouldBind(&statusTask); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Field cannot be blank",
		})
		return
	}

	err := appEnv.services.ChangeTaskStatus(statusTask.TaskID, statusTask.TaskStatus)
	if err != nil {
		if errors.Is(err, services.ErrUnprocessableEntity) {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": http.StatusText(http.StatusBadRequest),
			})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.Redirect(http.StatusSeeOther, "/")

}

// DeleteTheList deletes a specific to-do list.
func (appEnv *appEnv) DeleteTheList(c *gin.Context) {
	var deleteList DeleteListForm

	if err := c.ShouldBind(&deleteList); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Field cannot be blank",
		})
		return
	}

	err := appEnv.services.DeleteList(deleteList.ListID)
	if err != nil {
		if errors.Is(err, services.ErrNoRecordFound) {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": http.StatusText(http.StatusBadRequest),
			})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

// DeleteTheTask deletes a specific task.
func (appEnv *appEnv) DeleteTheTask(c *gin.Context) {
	var deleteTask DeleteTaskForm

	if err := c.ShouldBind(&deleteTask); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "Field cannot be blank",
		})
		return
	}

	err := appEnv.services.DeleteTask(deleteTask.TaskID)
	if err != nil {
		if errors.Is(err, services.ErrNoRecordFound) {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": http.StatusText(http.StatusBadRequest),
			})
		} else {
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.Redirect(http.StatusSeeOther, "/")

}
