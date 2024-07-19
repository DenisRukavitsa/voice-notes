package user

type UserModel struct {
	Id 			 string `bson:"_id"`
	Email    string
	Password string
}
