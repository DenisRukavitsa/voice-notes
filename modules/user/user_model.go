package user

type UserModel struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
