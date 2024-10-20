package domains

import (
	"context"
)

type AuthService interface {
	Auth(ctx context.Context, email string, password string) (HttpResponse, error)
}
