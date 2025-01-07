package complaint_service

import (
	"context"
	"errors"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"
	mbr "github.com/harrymuliawan03/layanan-pengaduan-api.git/app/repositories/complaint_repo"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/schemas"
	"gorm.io/gorm"
)

type ComplaintServiceImpl struct {
	complaintRepo mbr.ComplaintRepository
}

// FindAllByUserID implements ComplaintService.
func (c *ComplaintServiceImpl) FindAllByUserID(ctx context.Context, id uint) ([]dto.ComplaintData, error) {
	var result []dto.ComplaintData

	complaints, err := c.complaintRepo.FindAllByUserID(ctx, id)
	for _, complaint := range complaints {
		result = append(result, dto.ComplaintData{
			ID:          complaint.ID,
			Title:       complaint.Title,
			Description: complaint.Description,
			UserID:      complaint.UserID,
			Status:      complaint.Status,
			Photo:       complaint.Photo,
			Address:     complaint.Address,
	})
	
	}

	return result, err
}

// FindAllByUserID implements ComplaintService.
func (c *ComplaintServiceImpl) FindAll(ctx context.Context) ([]dto.ComplaintData, error) {
	var result []dto.ComplaintData

	complaints, err := c.complaintRepo.FindAll(ctx)
	for _, complaint := range complaints {
		result = append(result, dto.ComplaintData{
			ID:          complaint.ID,
			Title:       complaint.Title,
			Description: complaint.Description,
			UserID:      complaint.UserID,
			Status:      complaint.Status,
			Photo:       complaint.Photo,
			Address:     complaint.Address,
		})
	}

	return result, err
}

func NewComplaintService(cr mbr.ComplaintRepository) ComplaintService {
	return &ComplaintServiceImpl{complaintRepo: cr}
}

// Create implements ComplaintService.
func (c *ComplaintServiceImpl) Create(ctx context.Context, req requests.ComplaintCreateRequest) error {
	complaint := models.Complaint{UserID: req.UserID, Title: req.Title, Description: req.Description, Status: req.Status, Photo: req.Photo, Address: req.Address}
	return c.complaintRepo.Save(ctx, &complaint)
}

// Delete implements ComplaintService.
func (c *ComplaintServiceImpl) Delete(ctx context.Context, id uint) error {
	return c.complaintRepo.Delete(ctx, id)
}

// FindByID implements ComplaintService.
func (c *ComplaintServiceImpl) FindByID(ctx context.Context, id uint) (*dto.ComplaintData, error) {
	complaint, err := c.complaintRepo.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Complaint not found",
		}
		return nil, err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return nil, err
	}

	return &dto.ComplaintData{ID: complaint.ID, UserID: complaint.UserID, Title: complaint.Title, Description: complaint.Description, Status: complaint.Status, Photo: complaint.Photo, Address: complaint.Address}, nil
}

// Show implements ComplaintService.
func (c *ComplaintServiceImpl) Show(ctx context.Context, id uint) (*dto.ComplaintData, error) {
	panic("unimplemented")
}

// Update implements ComplaintService.
func (c *ComplaintServiceImpl) Update(ctx context.Context, id uint, req requests.ComplaintUpdateRequest) error {
	_, err := c.complaintRepo.FindByID(ctx, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorNotFound,
			Message: "Complaint not found",
		}
		return err
	} else if err != nil {
		err = &schemas.ResponseApiError{
			Status:  schemas.ApiErrorInternalServer,
			Message: err.Error(),
		}
		return err
	}

	complaint := models.Complaint{ID: id, Title: req.Title, Description: req.Description, Status: req.Status, Photo: req.Photo, Address: req.Address}
	return c.complaintRepo.Update(ctx, &complaint)
}
