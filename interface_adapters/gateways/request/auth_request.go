package request

type LoginRequest struct {
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required"`
}
