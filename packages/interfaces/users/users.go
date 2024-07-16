package users

type Users interface {
	Create()
	FindByEmail(email string)
}
