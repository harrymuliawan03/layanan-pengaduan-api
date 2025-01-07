package complaint_handler

import (
	"context"
	"net/http"
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/configs"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	cs "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/services/complaint_service"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/utils"

	"github.com/gofiber/fiber/v2"
)

type ComplaintHandler struct {
	cs  cs.ComplaintService
	cnf *configs.Config
}

func NewComplaintHandler(cs cs.ComplaintService, config *configs.Config) *ComplaintHandler {
	return &ComplaintHandler{
		cs:  cs,
		cnf: config,
	}
}

func (u *ComplaintHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req requests.ComplaintCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.cs.Create(c, req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiCreated(ctx, "Complaints created successfully", nil)
}

func (u *ComplaintHandler) FindAll(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := u.cs.FindAll(c)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get Complaints successfully", res)
}

func (u *ComplaintHandler) FindAllByUserID(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "Complaint id is required", http.StatusBadRequest, nil)
	}

	res, err := u.cs.FindAllByUserID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get Complaints successfully", res)
}

func (u *ComplaintHandler) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "Complaint id is required", http.StatusBadRequest, nil)
	}

	res, err := u.cs.FindByID(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "get Complaints successfully", res)
}

func (u *ComplaintHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")

	if id == 0 {
		return dto.ResponseApiError(ctx, "Complaint id is required", http.StatusBadRequest, nil)
	}

	var req requests.ComplaintUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusBadRequest, err.Error())
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		return dto.ResponseApiError(ctx, "validation error", http.StatusBadRequest, fails)
	}

	err := u.cs.Update(c, uint(id), req)
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "Complaints updated successfully", nil)
}

func (u *ComplaintHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id, _ := ctx.ParamsInt("id")
	if id == 0 {
		return dto.ResponseApiError(ctx, "Complaint id is required", http.StatusBadRequest, nil)
	}

	err := u.cs.Delete(c, uint(id))
	if err != nil {
		return dto.ResponseApiError(ctx, err.Error(), http.StatusInternalServerError, err.Error())
	}

	return dto.ResponseApiOk(ctx, "Complaints deleted successfully", nil)
}
