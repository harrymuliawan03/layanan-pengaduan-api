package complaint_update_handler

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	cur "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/complaint_update_repo"
	cus "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/complaint_update_service"

	"github.com/gofiber/fiber/v2"
)

func ComplaintUpdateRoute(route fiber.Router, cnf *configs.Config) {
	repo := cur.NewComplaintUpdateRepository()

	service := cus.NewComplaintUpdateService(repo)

	handler := NewComplaintUpdateHandler(service, cnf)
	complaintUpdate := route.Group("/complaint_updates")

	complaintUpdate.Get("/", handler.FindAll)
	complaintUpdate.Get("", handler.FindAll)
	complaintUpdate.Get("/:id", handler.Show)
	complaintUpdate.Put("/:id", handler.Update)
	complaintUpdate.Post("/", handler.Create)
	complaintUpdate.Post("", handler.Create)
	complaintUpdate.Delete("/:id", handler.Delete)
}
