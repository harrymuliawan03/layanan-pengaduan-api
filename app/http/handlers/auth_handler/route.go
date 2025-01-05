package authhandler

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	authverifrepo "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/auth_verif_repo"
	userrepo "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/user_repo"
	authservice "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/auth_service"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router, jwt fiber.Handler, cnf *configs.Config) {
	usr := userrepo.NewUserRepository()
	avr := authverifrepo.NewAuthVerifRepository()

	uss := authservice.NewAuthService(usr, cnf, avr)

	handler := NewUserHandler(uss, cnf)
	user := route.Group("/auth")

	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
}
