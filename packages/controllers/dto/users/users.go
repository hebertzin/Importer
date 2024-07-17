package users

type CreateUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"username"`
}
