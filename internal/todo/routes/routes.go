package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-golang/todo-app/internal/todo/handler"
	"github.com/m-golang/todo-app/internal/todo/helpers"
	"github.com/m-golang/todo-app/internal/todo/middleware"
)

// Router function sets up the routes for the to-do app.
// This function configures middleware and routes, and associates handlers to specific endpoints.
func Router() *gin.Engine {
	// Create a new Gin router for handling HTTP requests
	router := gin.Default()
	// Apply middleware globally
	// RecoverPanic ensures the app doesn't crash due to unhandled panics.
	router.Use(middleware.RecoverPanic())
	// SecureHeaders adds security-related headers to every response.
	router.Use(middleware.SecureHeaders())

	// Overrides the default 404 handler to return a custom JSON response
	// with a "Page not found" message and a 404 status when a route is not found.
	router.NoRoute(helpers.NotFoundPage)

	// Routes for to-do list and task management

	// GET / - Retrieve all to-do lists with their associated tasks
	// Maps to the handler function that fetches and displays all to-do lists and tasks.
	router.GET("/", handler.AllTodos) // Route to display all to-do lists

	// POST /list/new - Create a new to-do list
	// Maps to the handler function that creates a new to-do list based on the provided data.
	router.POST("/list/new", handler.NewList) // Route to create a new list

	// POST /task/new - Create a new task within a specified to-do list
	// Maps to the handler function that creates a new task for a specific to-do list.
	router.POST("/task/new", handler.NewTask) // Route to create a new task

	// PUT /task/status - Change the status of a task (e.g., mark as completed)
	// Maps to the handler function that updates the completion status of a task.
	router.PUT("/task/status", handler.TaskStatusChange) // Route to change the status of a task

	// DELETE /list/remove - Delete a specific to-do list
	// Maps to the handler function that deletes a to-do list by its ID.
	router.DELETE("/list/remove", handler.DeleteTheList) // Route to delete a specific list

	// DELETE /task/remove - Delete a specific task from a list
	// Maps to the handler function that deletes a task by its ID.
	router.DELETE("/task/remove", handler.DeleteTheTask) // Route to delete a specific task

	return router
}
