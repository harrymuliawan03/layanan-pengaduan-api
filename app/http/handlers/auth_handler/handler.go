package authhandler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/helpers"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/schemas"
	authservice "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/auth_service"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/color"
)

type UserHandler struct {
	authService authservice.AuthService
	cnf         *configs.Config
}

func NewUserHandler(authService authservice.AuthService, config *configs.Config) *UserHandler {
	return &UserHandler{
		authService: authService,
		cnf:         config,
	}
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.authService.Register(c, &req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "user created successfully", nil)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	res, err := u.authService.Login(c, &req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "user logged in successfully", res)
}

func (u *UserHandler) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var err error

	// Get the Authorization header
	authHeader := ctx.Get("Authorization")

	// Check if the token is present
	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		err = errors.New("missing or malformed JWT token")
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	// Extract the JWT token from the header
	tokenStr := authHeader[7:]

	// Parse the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.cnf.Jwt.Key), nil
	})

	if err != nil || !token.Valid {
		err = errors.New("invalid or expired token")
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())

	}

	// Extract claims and get the user ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New("unable to extract claims from token")
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())

	}

	color.Greenln("claims: ", claims["id"])
	userIDFloat, ok := claims["id"].(float64) // JWT stores numeric values as float64
	if !ok {
		err = errors.New("invalid token payload")
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}
	userID := uint(userIDFloat)

	res, err := u.authService.Show(c, userID)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "user retrieved successfully", res)
}

func (u *UserHandler) VerifEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	if ctx.Params("code") == "" {
		return dto.ResponseApiError(ctx, "missing verification code", http.StatusBadRequest, nil)
	}

	var req requests.VerifEmailRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.authService.VerifEmail(c, req, ctx.Params("code"))

	if err != nil {
		statuscode := helpers.CatchErrorResponseApi(err.(*schemas.ResponseApiError)).StatusCode
		return dto.ResponseApiError(ctx, err.Error(), statuscode, err.Error())
	}



	return dto.ResponseApiOk(ctx, "user email verified successfully", nil)
}
