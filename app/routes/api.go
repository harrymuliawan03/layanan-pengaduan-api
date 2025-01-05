package routes

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	authhandler "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/handlers/auth_handler"
	complaint_handler "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/handlers/complaint_handler"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/handlers/complaint_update_handler"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, jwt fiber.Handler, cnf *configs.Config) {
	api := app.Group("/api")
	authhandler.UserRoute(api, jwt, cnf)
	complaint_handler.ComplaintRoute(api, cnf)
	complaint_update_handler.ComplaintUpdateRoute(api, cnf)

	// Listing routes
	pkg.ListRoutes(app)
}
