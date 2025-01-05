package main

import (
	"log"
	"os"
	"strconv"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/middleware"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/routes"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/cmd/cli/commands"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gookit/color"
)

func main() {
	cnf := configs.Load()
	configs.ConnectDB()
	if len(os.Args) >= 2 {
		commands.Execute()
		return
	}

	app := fiber.New(
		fiber.Config{
			CaseSensitive:         true,
			StrictRouting:         true,
			AppName:               cnf.App.Name,
			DisableStartupMessage: true,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				if err == fiber.ErrInternalServerError {
					return dto.ResponseApiError(ctx, err.Error(), 500, nil)
				}
				return dto.ResponseApiError(ctx, err.Error(), 500, nil)
			},
		})

	// Add CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	midd := middleware.NewAuthMMiddleware(cnf.Jwt.Key)
	routes.RegisterRoutes(app, midd, cnf)
	appPort := strconv.Itoa(cnf.App.Port)
	color.Blueln("This app is running on 127.0.0.1:" + appPort)
	log.Fatal(app.Listen(":" + appPort))
}
