package response

import "undefeated-davout/echo-api-sample/entities"

type TaskResponse struct {
	ID     entities.TaskID     `json:"id"`
	Title  string              `json:"title"`
	Status entities.TaskStatus `json:"status"`
}

type ListTasksResponse []TaskResponse

type AddTaskResponse struct {
	ID entities.TaskID `json:"id"`
}
