package request

type AddTaskRequest struct {
	Title  string `form:"title" validate:"required"`
	Status string `form:"status" validate:"required"`
}
