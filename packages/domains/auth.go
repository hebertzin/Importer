package domains

import (
	"context"
)

type AuthUseCase interface {
	Auth(ctx context.Context, email string, password string) (HttpResponse, error)
}
