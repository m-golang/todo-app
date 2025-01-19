package forms

// NewListForm is used to create a new to-do list.
type NewListForm struct {
	Name string `json:"new_list" binding:"required"`
}

// NewTaskForm is used to create a new task.
type NewTaskFrom struct {
	Task   string `json:"new_task" binding:"required"`
	ListID int    `json:"list_id" binding:"required"`
}

// TaskStatusForm is used to update the status of a task.
type TaskStatusForm struct {
	TaskID     int `json:"id_task" binding:"required"`
	TaskStatus string `json:"status_task" binding:"required"`
}

// DeleteListForm is used to delete a to-do list.
type DeleteListForm struct {
	ListID int `json:"list_id" binding:"required"`
}

// DeleteTaskForm is used to delete a task.
type DeleteTaskForm struct {
	TaskID int `json:"id_task" binding:"required"`
}
