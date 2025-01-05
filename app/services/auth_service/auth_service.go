package authservice

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
)

type AuthService interface {
	Show(ctx context.Context, id uint) (*dto.UserData, error)
	Register(ctx context.Context, req *requests.RegisterRequest) error
	Login(ctx context.Context, req *requests.LoginRequest) (dto.LoginResponse, error)
	Update(ctx context.Context, req requests.UserUpdateRequest) error
	Delete(ctx context.Context, id uint) error
	VerifEmail(ctx context.Context, req requests.VerifEmailRequest, code string) error
}
