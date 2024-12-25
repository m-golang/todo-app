package main

import "github.com/gin-gonic/gin"

// Router function sets up the routes for the to-do app
func (appEnv *appEnv) Router(router *gin.Engine) {

	router.GET("/", appEnv.AllTodos) // Route to display all to-do lists

	router.POST("/list/new", appEnv.NewList) // Route to create a new list
	router.POST("/task/new", appEnv.NewTask) // Route to create a new task

	router.POST("/task/status", appEnv.TaskStatusChange) // Route to change the status of a task

	router.POST("/list/remove", appEnv.DeleteTheList) // Route to delete a specific list
	router.POST("/task/remove", appEnv.DeleteTheTask) // Route to delete a specific task

}
