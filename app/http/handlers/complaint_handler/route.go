package complaint_handler

import (
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	cr "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/complaint_repo"
	cs "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/complaint_service"

	"github.com/gofiber/fiber/v2"
)

func ComplaintRoute(route fiber.Router, cnf *configs.Config) {
	repo := cr.NewComplaintRepository()

	service := cs.NewComplaintService(repo)

	handler := NewComplaintHandler(service, cnf)
	complaint := route.Group("/complaints")

	complaint.Get("/", handler.FindAll)
	complaint.Get("", handler.FindAll)
	complaint.Get("/:id", handler.Show)
	complaint.Put("/:id", handler.Update)
	complaint.Post("/", handler.Create)
	complaint.Post("", handler.Create)
	complaint.Delete("/:id", handler.Delete)
}
