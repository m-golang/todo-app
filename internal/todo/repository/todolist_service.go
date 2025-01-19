package repository

import "fmt"

// FetchListsWithTasks retrieves all to-do lists and their associated tasks
// from the service layer, groups the tasks by list ID, and returns the lists
// with their corresponding tasks. This function handles errors in retrieving
// the lists and tasks, ensuring that appropriate error messages are returned.
//
// It is the responsibility of the caller to handle the returned data and
// any potential errors.
//
// Returns:
// - A slice of TodoList pointers with tasks populated, or
// - An error if there was an issue fetching the lists or tasks.
func FetchListsWithTasks() ([]*TodoList, error) {
	// Fetch all to-do lists from the service layer
	allLists, err := RetrieveAllLists()
	if err != nil {
		// Return a wrapped error if fetching lists fails
		return nil, fmt.Errorf("failed to retrieve todos data: %w", err)

	}

	// Fetch all tasks from the service layer
	allTasks, err := RetrieveAllTasks()
	if err != nil {
		// Return a wrapped error if fetching tasks fails
		return nil, fmt.Errorf("failed to retrieve todos data: %w", err)

	}

	// Create a map to group tasks by their associated to-do list ID
	tasksByList := make(map[int][]*TodoTasks)

	// Iterate over all tasks and group them by the list they belong to
	for _, task := range allTasks {
		tasksByList[task.Id_list_task] = append(tasksByList[task.Id_list_task], task)
	}

	// Iterate over all to-do lists and assign the corresponding tasks from the 'tasksByList' map
	for i := range allLists {
		listID := allLists[i].List_id
		allLists[i].Tasks = tasksByList[listID]
	}

	// Return the lists with the associated tasks
	return allLists, nil
}
