package domain

type Token interface {
	SignIn(email string) (string, error)
	Verify(tokenString string) (Claims, error)
}

type Claims struct {
	Email string
}
