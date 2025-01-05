package complaint_update_service

import (
	"context"
	"errors"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"
	cus "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/complaint_update_repo"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/schemas"
	"gorm.io/gorm"
)

type ComplaintUpdateServiceImpl struct {
	complaintUpdateRepo cus.ComplaintUpdateRepository
}

// FindAllByUserID implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) FindAll(ctx context.Context) ([]dto.ComplaintUpdateData, error) {
	var result []dto.ComplaintUpdateData

	complaintUpdates, err := c.complaintUpdateRepo.FindAll(ctx)
	for _, complaintUpdate := range complaintUpdates {
		result = append(result, dto.ComplaintUpdateData{
			ID:          complaintUpdate.ID,
			AdminID:     complaintUpdate.AdminID,
			ComplaintID: complaintUpdate.ComplaintID,
			Status:      complaintUpdate.Status,
			Note:        complaintUpdate.Note,
			UpdatedAt:   complaintUpdate.UpdatedAt,
		})
	}

	return result, err
}

func NewComplaintUpdateService(cr cus.ComplaintUpdateRepository) ComplaintUpdateService {
	return &ComplaintUpdateServiceImpl{complaintUpdateRepo: cr}
}

// Create implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) Create(ctx context.Context, req requests.ComplaintUpdateCreateRequest) error {
	complaintUpdate := models.ComplaintUpdate{AdminID: req.AdminID, ComplaintID: req.ComplaintID, Status: req.Status, Note: req.Note}
	return c.complaintUpdateRepo.Save(ctx, &complaintUpdate)
}

// Delete implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) Delete(ctx context.Context, id uint) error {
	return c.complaintUpdateRepo.Delete(ctx, id)
}

// FindByID implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) FindByID(ctx context.Context, id uint) (*dto.ComplaintUpdateData, error) {
	complaintUpdate, err := c.complaintUpdateRepo.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "ComplaintUpdate not found",
		}
		return nil, err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return nil, err
	}

	return &dto.ComplaintUpdateData{ID: complaintUpdate.ID, ComplaintID: complaintUpdate.ComplaintID, AdminID: complaintUpdate.AdminID, Status: complaintUpdate.Status, Note: complaintUpdate.Note, UpdatedAt: complaintUpdate.UpdatedAt}, nil
}

// Show implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) Show(ctx context.Context, id uint) (*dto.ComplaintUpdateData, error) {
	panic("unimplemented")
}

// Update implements ComplaintUpdateService.
func (c *ComplaintUpdateServiceImpl) Update(ctx context.Context, id uint, req requests.ComplaintUpdateUpdateRequest) error {
	_, err := c.complaintUpdateRepo.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "ComplaintUpdate not found",
		}
		return err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return err
	}

	complaintUpdate := models.ComplaintUpdate{ID: id, AdminID: req.AdminID, ComplaintID: req.ComplaintID, Status: req.Status, Note: req.Note}
	return c.complaintUpdateRepo.Update(ctx, &complaintUpdate)
}
