package schema

type User struct {
	ID       string `form:"id"`
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	FullName string `form:"full_name" validate:"required"`
}
