package schema

type RefreshToken struct {
	ID     string `form:"id" validate:"required"`
	Token  string `form:"token" validate:"required"`
	UserID string `form:"user_id" validate:"required"`
}
