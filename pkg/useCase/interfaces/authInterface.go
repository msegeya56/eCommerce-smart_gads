package interfaces

import (
	"context"

	"github.com/msegeya56/eCommerce-smart_gads/pkg/domain"
)

type AuthService interface {
	SignUp(ctx context.Context, user domain.Users) error
	Login(ctx context.Context, user domain.Users) (domain.Users, error)
	OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error)
}
