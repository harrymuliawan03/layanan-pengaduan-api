package middleware

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == fiber.ErrInternalServerError {
				return dto.ResponseApiError(ctx, err.Error(), 500, nil)
			}
			return dto.ResponseApiError(ctx, "Unauthorized", 500, nil)
		},
	})
}
