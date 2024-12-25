package main

// NewListForm is used to create a new to-do list.
type NewListForm struct {
	Name string `form:"New_list" json:"New_list" binding:"required"`
}

// NewTaskForm is used to create a new task.
type NewTaskFrom struct {
	Task   string `form:"New_task" json:"New_task" binding:"required"`
	ListID int    `form:"List_id" json:"List_id" binding:"required"`
}

// TaskStatusForm is used to update the status of a task.
type TaskStatusForm struct {
	TaskID     string    `form:"Id_task" json:"Id_task" binding:"required"`
	TaskStatus string `form:"Status_task" json:"Status_task" binding:"required"`
}

// DeleteListForm is used to delete a to-do list.
type DeleteListForm struct {
	ListID int `form:"List_id" json:"List_id" binding:"required"`
}

// DeleteTaskForm is used to delete a task.
type DeleteTaskForm struct {
	TaskID int `form:"Id_task" json:"Id_task" binding:"required"`
}
