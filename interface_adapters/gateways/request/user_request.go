package request

type AddUserRequest struct {
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required"`
	Role     string `form:"role" validate:"required"`
}
