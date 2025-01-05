package complaint_update_handler

import (
	"context"
	"net/http"
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	cus "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/complaint_update_service"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/utils"

	"github.com/gofiber/fiber/v2"
)

type ComplaintUpdateHandler struct {
	cus cus.ComplaintUpdateService
	cnf *configs.Config
}

func NewComplaintUpdateHandler(cus cus.ComplaintUpdateService, config *configs.Config) *ComplaintUpdateHandler {
	return &ComplaintUpdateHandler{
		cus: cus,
		cnf: config,
	}
}

func (u *ComplaintUpdateHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.ComplaintUpdateCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.cus.Create(c, req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiCreated(ctx, "ComplaintUpdates created successfully", nil)
}

func (u *ComplaintUpdateHandler) FindAll(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := u.cus.FindAll(c)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get ComplaintUpdates successfully", res)
}

func (u *ComplaintUpdateHandler) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "ComplaintUpdate id is required", http.StatusBadRequest, nil)
	}

	res, err := u.cus.FindByID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get ComplaintUpdates successfully", res)
}

func (u *ComplaintUpdateHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "ComplaintUpdate id is required", http.StatusBadRequest, nil)
	}

	var req requests.ComplaintUpdateUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.cus.Update(c, uint(id), req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "ComplaintUpdates updated successfully", nil)
}

func (u *ComplaintUpdateHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		return dto.ResponseApiError(ctx, "ComplaintUpdate id is required", http.StatusBadRequest, nil)
	}

	err := u.cus.Delete(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "ComplaintUpdates deleted successfully", nil)
}
