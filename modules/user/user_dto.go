package user

type UserDto struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
